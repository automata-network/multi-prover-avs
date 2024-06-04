package xtask

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	// "scroll-tech/common/types/encoding/codev1"
	// "scroll-tech/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/automata-network/multi-prover-avs/xmetric"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type TaskType int

const (
	MinTaskType TaskType = 0
	ScrollTask  TaskType = 1
	MaxTaskType TaskType = 2
)

func (t TaskType) Value() string {
	switch t {
	case ScrollTask:
		return "scroll"
	default:
		return fmt.Sprint(int(t))
	}
}

var (
	ErrInvalidTaskManager = logex.Define("invalid task type: %v")
)

type TaskManagerConfig struct {
	Identifier TaskType

	Prover     string
	L2Endpoint string

	Endpoint  string
	Topics    [][]common.Hash
	Addresses []common.Address
	// OffsetFile       string
	ScanIntervalSecs int64
}

type TaskContext struct {
	Prover    *ProverClient
	EthClient *rpc.Client
}

type TaskManager struct {
	sampling int64
	sources  map[string]*ethclient.Client
	tracers  map[TaskType]*utils.LogTracer
	contexts map[TaskType]*TaskContext

	collector       *xmetric.AggregatorCollector
	referenceClient eth.Client

	tasksMutex       sync.Mutex
	tasks            map[TaskType]*TaskTuple
	presetStartBlock uint64
}

type TaskTuple struct {
	Info     *TaskInfo
	Channels []chan struct{}
}

type TaskInfo struct {
	Type    TaskType
	Context json.RawMessage
	TaskID  int64
	Ext     json.RawMessage
}

func NewTaskManager(collector *xmetric.AggregatorCollector, sampling int64, referenceClient eth.Client, tasks []*TaskManagerConfig) (*TaskManager, error) {
	sources := make(map[string]*ethclient.Client)
	tracers := make(map[TaskType]*utils.LogTracer)
	contexts := make(map[TaskType]*TaskContext)
	tm := &TaskManager{
		sampling:         sampling,
		sources:          sources,
		tracers:          tracers,
		contexts:         contexts,
		collector:        collector,
		referenceClient:  referenceClient,
		presetStartBlock: 0, //19976077,
		tasks:            make(map[TaskType]*TaskTuple, MaxTaskType),
	}

	for _, cfg := range tasks {
		if cfg.Identifier >= MaxTaskType || cfg.Identifier <= MinTaskType {
			return nil, ErrInvalidTaskManager.Format(cfg.Identifier)
		}

		var err error
		taskCtx := &TaskContext{}
		contexts[cfg.Identifier] = taskCtx
		if cfg.Prover != "" {
			taskCtx.Prover, err = NewProverClient(cfg.Prover)
			if err != nil {
				return nil, logex.Trace(err, fmt.Sprintf("connect to prover: %v", cfg.Prover))
			}
		}
		if cfg.L2Endpoint != "" {
			taskCtx.EthClient, err = rpc.Dial(cfg.L2Endpoint)
			if err != nil {
				return nil, logex.Trace(err, fmt.Sprintf("connect to l2endpoint: %v", cfg.L2Endpoint))
			}
		}

		source, ok := sources[cfg.Endpoint]
		if !ok {
			source, err = ethclient.Dial(cfg.Endpoint)
			if err != nil {
				return nil, logex.Trace(err, fmt.Sprintf("dial taskFetcher.Endpoint: %q", cfg.Endpoint))
			}
			sources[cfg.Endpoint] = source
		}

		if _, ok := tracers[cfg.Identifier]; ok {
			return nil, logex.NewErrorf("invalid task: %v is already exists", cfg.Identifier)
		}

		tracers[cfg.Identifier] = utils.NewLogTracer(source, &utils.LogTracerConfig{
			Id:               fmt.Sprintf("aggregator-task-fetcher-%v", cfg.Identifier),
			Wait:             5,
			Max:              100,
			Topics:           cfg.Topics,
			Addresses:        cfg.Addresses,
			ScanIntervalSecs: cfg.ScanIntervalSecs,
			SkipOnError:      true,
			Handler:          tm,
		})
	}

	return tm, nil
}

func (t *TaskManager) OnNewLog(ctx context.Context, log *types.Log) error {
	source := utils.KeyLogTracerSourceClient{}.Get(ctx)
	id := ctx.Value(TaskManagerId{}).(TaskType)
	t.collector.NewTask.WithLabelValues(id.Value()).Add(1)

	switch id {
	case ScrollTask:
		if err := t.onScrollTask(ctx, source, log); err != nil {
			return logex.Trace(err)
		}
	default:
		return nil
	}
	return nil
}

func (t *TaskManager) GetNextTask(ctx context.Context, ty TaskType, withContext bool, prevTaskID int64, timeout time.Duration) (*TaskInfo, bool) {
	t.collector.FetchTask.WithLabelValues(ty.Value(), fmt.Sprint(withContext)).Add(1)

	var waitChan *chan struct{}
	timer := time.NewTimer(timeout)
	wait := func() bool {
		select {
		case <-timer.C:
			return false
		case <-ctx.Done():
			return false
		case <-*waitChan:
			return true
		}
	}
	retry := true

	for {
		t.tasksMutex.Lock()
		taskTuple := t.tasks[ty]
		if taskTuple == nil {
			taskTuple = new(TaskTuple)
			t.tasks[ty] = taskTuple
		}
		skip := taskTuple.Info == nil
		skip = skip || taskTuple.Info.Context == nil && withContext
		skip = skip || taskTuple.Info.TaskID <= prevTaskID
		if skip {
			ch := make(chan struct{})
			waitChan = &ch
			taskTuple.Channels = append(taskTuple.Channels, ch)
		}
		t.tasksMutex.Unlock()
		if skip {
			if !retry {
				return nil, false
			}
			retry = wait()
			continue
		}
		if !withContext && len(taskTuple.Info.Context) != 0 {
			newInfo := *taskTuple.Info
			newInfo.Context = nil
			return &newInfo, true
		}
		return taskTuple.Info, true
	}
}

func (t *TaskManager) updateTask(taskInfo TaskInfo) {
	t.tasksMutex.Lock()
	taskTuple := t.tasks[taskInfo.Type]
	if taskTuple == nil {
		taskTuple = new(TaskTuple)
		t.tasks[taskInfo.Type] = taskTuple
	}
	taskTuple.Info = &taskInfo
	chs := taskTuple.Channels
	taskTuple.Channels = nil
	t.tasksMutex.Unlock()
	for _, ch := range chs {
		close(ch)
	}
}

func (t *TaskManager) onScrollTask(ctx context.Context, source *ethclient.Client, log *types.Log) error {
	prover := ctx.Value(TaskManagerProverClient{}).(*ProverClient)
	referenceBlockNumber, err := t.referenceClient.BlockNumber(ctx)
	if err != nil {
		return logex.Trace(err)
	}

	batchId := new(big.Int).SetBytes(log.Topics[1][:])
	if t.sampling > 0 {
		if batchId.Int64()%t.sampling != 0 {
			logex.Infof("scroll: sampling miss for batchId#%v", batchId)
			return nil
		}
	}

	taskInfo := &TaskInfo{
		Type:   ScrollTask,
		TaskID: batchId.Int64(),
	}

	tx, _, err := source.TransactionByHash(ctx, log.TxHash)
	if err != nil {
		return logex.Trace(err)
	}

	args, err := ScrollABI.Methods["commitBatch"].Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		return logex.Trace(err)
	}

	startBlock := int64(0)
	endBlock := int64(0)
	for _, chunk := range args[2].([][]byte) {
		for i := 0; i < int(chunk[0]); i++ {
			blockNumber := int64(binary.BigEndian.Uint64(chunk[1:][i*60 : i*60+8]))
			if startBlock == 0 {
				startBlock = blockNumber
			} else {
				endBlock = blockNumber
			}
		}
	}

	taskInfo.Ext, err = json.Marshal(ScrollTaskExt{
		StartBlock:           (*hexutil.Big)(big.NewInt(startBlock)),
		EndBlock:             (*hexutil.Big)(big.NewInt(endBlock)),
		BatchData:            (hexutil.Bytes)(tx.Data()[4:]),
		CommitTx:             log.TxHash,
		ReferenceBlockNumber: referenceBlockNumber - 1,
	})
	if err != nil {
		return logex.Trace(err)
	}

	t.updateTask(*taskInfo)

	startGenerateContext := time.Now()
	taskCtx, ignore, err := prover.GenerateContext(ctx, startBlock, endBlock, taskInfo.Type)
	if ignore {
		return nil
	}
	if err != nil {
		return logex.Trace(err, fmt.Sprintf("fetching context for scroll batchId#%v", batchId))
	}
	generateContextCost := time.Since(startGenerateContext).Truncate(time.Millisecond)

	// logex.Info(prover.GetPoeByPob(ctx, tx.Data()[4:], taskCtx))

	taskInfo.Context, err = json.Marshal(taskCtx)
	if err != nil {
		return logex.Trace(err)
	}
	t.updateTask(*taskInfo)
	logex.Infof("update task: [%v] %v, (generateContext:%v)", taskInfo.Type.Value(), taskInfo.TaskID, generateContextCost)
	time.Sleep(10 * time.Second)

	return nil
}

func (t *TaskManager) SaveBlock(uint64) error {
	return nil
}

func (t *TaskManager) GetBlock() (uint64, error) {
	return t.presetStartBlock, nil
}

func (t *TaskManager) Run(ctx context.Context) error {
	errChan := make(chan error, len(t.tracers))
	var wg sync.WaitGroup
	for id := range t.tracers {
		wg.Add(1)
		tracer := t.tracers[id]
		taskCtx := t.contexts[id]
		ctx := context.WithValue(ctx, TaskManagerId{}, id)
		ctx = context.WithValue(ctx, TaskManagerProverClient{}, taskCtx.Prover)
		ctx = context.WithValue(ctx, TaskManagerEthClient{}, taskCtx.EthClient)
		go func() {
			defer wg.Done()
			if err := tracer.Run(ctx); err != nil {
				errChan <- logex.Trace(err)
			}
		}()
	}
	wg.Wait()
	close(errChan)
	return <-errChan
}

type TaskManagerId struct{}
type TaskManagerProverClient struct{}
type TaskManagerEthClient struct{}
