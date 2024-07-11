package aggregator

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/automata-network/multi-prover-avs/xmetric"
	"github.com/automata-network/multi-prover-avs/xtask"
	"github.com/chzyer/logex"
)

type AggregatorApi struct {
	agg *Aggregator
}

type SignedTaskResponse struct {
	TaskResponse MultiProverServiceManager.IMultiProverServiceManagerStateHeader
	BlsSignature bls.Signature
	OperatorId   types.OperatorId
}

type FetchTaskReq struct {
	PrevTaskID  int            `json:"prev_task_id"`
	TaskType    xtask.TaskType `json:"task_type"`
	MaxWaitSecs int            `json:"max_wait_secs"`
	WithContext bool           `json:"with_context"`
}

type FetchTaskResp struct {
	Ok       bool            `json:"ok"`
	TaskID   int             `json:"task_id"`
	TaskType xtask.TaskType  `json:"task_type"`
	Ext      json.RawMessage `json:"ext"`
	Context  json.RawMessage `json:"context,omitempty"`
}

func (a *AggregatorApi) FetchTask(ctx context.Context, req *FetchTaskReq) (*FetchTaskResp, error) {
	defer func() {
		if err := recover(); err != nil {
			logex.Error(err, string(debug.Stack()))
			panic(err)
		}
	}()

	rsp, err := a.fetchTask(ctx, req)
	if err != nil {
		logex.Error(err)
		return nil, err
	}
	return rsp, nil
}

func (a *AggregatorApi) fetchTask(ctx context.Context, req *FetchTaskReq) (*FetchTaskResp, error) {
	if req.MaxWaitSecs > 30 {
		req.MaxWaitSecs = 30
	}

	timeout := time.Duration(req.MaxWaitSecs) * time.Second
	taskInfo, ok := a.agg.TaskManager.GetNextTask(ctx, req.TaskType, req.WithContext, int64(req.PrevTaskID), timeout)
	if !ok {
		return &FetchTaskResp{
			Ok: false,
		}, nil
	}

	return &FetchTaskResp{
		Ok:       true,
		TaskID:   int(taskInfo.TaskID),
		TaskType: taskInfo.Type,
		Ext:      taskInfo.Ext,
		Context:  taskInfo.Context,
	}, nil
}

func (a *AggregatorApi) SubmitTask(ctx context.Context, req *TaskRequest) error {
	defer func() {
		if err := recover(); err != nil {
			logex.Pretty(req)
			logex.Error(err, string(debug.Stack()))
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := a.submitTask(ctx, req); err != nil {
		logex.Error(err)
		return nil
	}
	return nil
}

func (a *AggregatorApi) submitTask(ctx context.Context, req *TaskRequest) error {
	taskCtx := []string{fmt.Sprintf("id=%x", req.OperatorId)}
	start := time.Now()
	defer func() {
		logex.Infof("SubmitTask cost time: %v, ctx=%v", time.Since(start), taskCtx)
	}()
	// check bls public key
	digest, err := req.Task.Digest()
	if err != nil {
		return logex.Trace(err, taskCtx)
	}
	taskCtx = append(taskCtx, fmt.Sprintf("digest=%x", digest))

	operatorAddr, err := bindings.GetOperatorAddrByOperatorID(a.agg.client, a.agg.registryCoordinator, req.OperatorId)
	if err != nil {
		return logex.Trace(err, taskCtx)
	}
	taskCtx = append(taskCtx, fmt.Sprintf("addr=%v", operatorAddr))

	operatorPubkeys, err := a.agg.registryCache.GetOperatorsAvsStateAtBlock(ctx, utils.BytesToQuorumNums(req.Task.QuorumNumbers), req.Task.ReferenceBlockNumber)
	if err != nil {
		return logex.Trace(err, taskCtx)
	}
	pubkey, ok := operatorPubkeys[req.OperatorId]
	if !ok {
		return logex.NewErrorf("operatorId not registered, ctx=%v", taskCtx)
	}
	taskCtx = append(taskCtx, fmt.Sprintf("bls=%x", pubkey.Pubkeys.G1Pubkey.Serialize()))

	validPubkey, err := req.Signature.Verify(pubkey.Pubkeys.G2Pubkey, digest)
	if err != nil {
		return logex.Trace(err, taskCtx)
	}
	if !validPubkey {
		return logex.NewErrorf("invalid signature&pubkey, ctx=%v", taskCtx)
	}
	x, y := utils.SplitPubkey(pubkey.Pubkeys.G1Pubkey.Serialize())

	pass, err := a.agg.verifyKey(x, y)
	if err != nil {
		return logex.Trace(err, taskCtx)
	}
	if !pass {
		return logex.NewErrorf("prover not registered, ctx=%v", taskCtx)
	}

	if err := a.agg.submitStateHeader(ctx, req); err != nil {
		return logex.Trace(err, taskCtx)
	}

	logex.Infof("receive task: %v, ctx=%v", req.Task, taskCtx)
	return nil
}

type SubmitMetricsReq struct {
	Name    string                  `json:"name"`
	Metrics []*xmetric.MetricFamily `json:"metrics"`
}

func (a *AggregatorApi) SubmitMetrics(ctx context.Context, req *SubmitMetricsReq) error {
	logex.Infof("accept %v metrics from [%v]", len(req.Metrics), req.Name)
	a.agg.Collector.AddOperatorMetrics(req.Name, req.Metrics)
	return nil
}
