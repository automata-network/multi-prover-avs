package aggregator

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/automata-network/multi-prover-avs/contracts/bindings"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/RegistryCoordinator"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/TEELivenessVerifier"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/automata-network/multi-prover-avs/xmetric"
	"github.com/automata-network/multi-prover-avs/xtask"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Config struct {
	ListenAddr       string
	TimeToExpirySecs int

	EcdsaPrivateKey                      string
	EthHttpEndpoint                      string
	EthWsEndpoint                        string
	AttestationLayerRpcURL               string
	MultiProverContractAddress           common.Address
	TEELivenessVerifierContractAddressV1 common.Address
	TEELivenessVerifierContractAddress   common.Address

	AVSRegistryCoordinatorAddress common.Address
	OperatorStateRetrieverAddress common.Address
	EigenMetricsIpPortAddress     string
	ScanStartBlock                uint64
	Threshold                     uint64
	Sampling                      uint64

	OpenTelemetry *xmetric.OpenTelemetryConfig

	TaskFetcher []*xtask.TaskManagerConfig

	Simulation bool
}

type Aggregator struct {
	cfg *Config

	blsAggregationService BlsAggregationService
	transactOpt           *bind.TransactOpts

	TaskManager *xtask.TaskManager

	client *ethclient.Client

	multiProverContract   *MultiProverServiceManager.MultiProverServiceManager
	TEELivenessVerifierV1 *TEELivenessVerifier.TEELivenessVerifierCaller
	TEELivenessVerifierV2 *TEELivenessVerifier.TEELivenessVerifierCaller
	registry              *avsregistry.AvsRegistryServiceChainCaller
	registryCoordinator   *RegistryCoordinator.RegistryCoordinator

	eigenClients *clients.Clients

	Collector *xmetric.AggregatorCollector

	taskMutex    sync.Mutex
	taskIndexSeq uint32
	taskIndexMap map[types.TaskResponseDigest]*Task
}

type Task struct {
	state *TaskRequest
	index uint32
}

func NewAggregator(ctx context.Context, cfg *Config) (*Aggregator, error) {
	if cfg.Sampling == 0 {
		cfg.Sampling = 2000
	}
	logex.Info("Multi Prover Aggregator Initializing...")
	ecdsaPrivateKey, err := crypto.HexToECDSA(cfg.EcdsaPrivateKey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	client, err := ethclient.Dial(cfg.EthHttpEndpoint)
	if err != nil {
		return nil, logex.Trace(err, fmt.Sprintf("dial:%q", cfg.EthHttpEndpoint))
	}
	attestationClient, err := ethclient.Dial(cfg.AttestationLayerRpcURL)
	if err != nil {
		return nil, logex.Trace(err, fmt.Sprintf("connecting to AttestationLayerRpcURL:%q", cfg.AttestationLayerRpcURL))
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, logex.Trace(err, "fetch chainID")
	}
	transactOpt, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		return nil, logex.Trace(err)
	}
	logger := utils.NewLogger(logex.NewLoggerEx(os.Stderr))

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthHttpEndpoint,
		EthWsUrl:                   cfg.EthWsEndpoint,
		RegistryCoordinatorAddr:    cfg.AVSRegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: cfg.OperatorStateRetrieverAddress.String(),
		AvsName:                    "aggregator",
		PromMetricsIpPortAddress:   cfg.EigenMetricsIpPortAddress,
	}

	eigenClients, err := clients.BuildAll(chainioConfig, ecdsaPrivateKey, logger)
	if err != nil {
		return nil, logex.Trace(err)
	}

	multiProverContract, err := MultiProverServiceManager.NewMultiProverServiceManager(cfg.MultiProverContractAddress, client)
	if err != nil {
		return nil, logex.Trace(err)
	}
	teeLivenessVerifier, err := TEELivenessVerifier.NewTEELivenessVerifierCaller(cfg.TEELivenessVerifierContractAddress, attestationClient)
	if err != nil {
		return nil, logex.Trace(err)
	}
	var teeLivenessVerifierV1 *TEELivenessVerifier.TEELivenessVerifierCaller
	var emptyAddr common.Address
	if cfg.TEELivenessVerifierContractAddressV1 != emptyAddr {
		teeLivenessVerifierV1, err = TEELivenessVerifier.NewTEELivenessVerifierCaller(cfg.TEELivenessVerifierContractAddressV1, attestationClient)
		if err != nil {
			return nil, logex.Trace(err)
		}
	}

	collector := xmetric.NewAggregatorCollector("avs")

	taskManager, err := xtask.NewTaskManager(collector, int64(cfg.Sampling), eigenClients.EthHttpClient, cfg.TaskFetcher)
	if err != nil {
		return nil, logex.Trace(err)
	}

	operatorPubkeysService, err := NewOperatorPubkeysService(ctx, client, eigenClients.AvsRegistryChainSubscriber, eigenClients.AvsRegistryChainReader, logger, "", cfg.ScanStartBlock, 5000)
	if err != nil {
		return nil, logex.Trace(err)
	}
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(eigenClients.AvsRegistryChainReader, operatorPubkeysService, logger)
	blsAggregationService := NewBlsAggregatorService(avsRegistryService, logger)

	registryCoordinator, err := RegistryCoordinator.NewRegistryCoordinator(cfg.AVSRegistryCoordinatorAddress, client)
	if err != nil {
		return nil, logex.Trace(err)
	}

	return &Aggregator{
		cfg:                   cfg,
		transactOpt:           transactOpt,
		client:                client,
		eigenClients:          eigenClients,
		blsAggregationService: blsAggregationService,
		multiProverContract:   multiProverContract,
		TEELivenessVerifierV1: teeLivenessVerifierV1,
		TEELivenessVerifierV2: teeLivenessVerifier,
		registryCoordinator:   registryCoordinator,
		registry:              avsRegistryService,
		TaskManager:           taskManager,
		taskIndexMap:          make(map[types.Bytes32]*Task),
		Collector:             collector,
	}, nil
}

func (agg *Aggregator) startUpdateOperators(ctx context.Context) (func() error, error) {
	quorumNums := types.QuorumNums{0}
	blockNumber, err := agg.client.BlockNumber(ctx)
	if err != nil {
		return nil, logex.Trace(err)
	}
	states, err := agg.registry.GetOperatorsAvsStateAtBlock(ctx, quorumNums, uint32(blockNumber))
	if err != nil {
		return nil, logex.Trace(err)
	}
	var operators []common.Address
	for k := range states {
		operatorAddr, err := agg.eigenClients.AvsRegistryChainReader.GetOperatorFromId(nil, k)
		if err != nil {
			return nil, logex.Trace(err)
		}
		isRegistered, err := agg.eigenClients.AvsRegistryChainReader.IsOperatorRegistered(nil, operatorAddr)
		if err != nil {
			return nil, logex.Trace(err)
		}
		if isRegistered {
			operators = append(operators, operatorAddr)
		}
	}

	newOpt := *agg.transactOpt
	newOpt.NoSend = true
	for i := 1; i < len(operators); i++ {
		tx, err := agg.registryCoordinator.UpdateOperators(&newOpt, operators[:i])
		if err != nil {
			return nil, logex.Trace(err)
		}
		logex.Infof("tx hash: %v -> %v", i, tx.Gas())
	}
	// logex.Info(states)
	return func() error { return nil }, nil
}

func (agg *Aggregator) verifyKey(x [32]byte, y [32]byte) (bool, error) {
	if agg.TEELivenessVerifierV1 != nil {
		pass, err := agg.TEELivenessVerifierV1.VerifyLivenessProof(nil, x, y)
		if err != nil {
			return false, logex.Trace(err, "v1")
		}
		if pass {
			return true, nil
		}
	}

	pass, err := agg.TEELivenessVerifierV2.VerifyLivenessProof(nil, x, y)
	if err != nil {
		return false, logex.Trace(err, "v2")
	}
	return pass, nil
}

func (agg *Aggregator) startRpcServer(ctx context.Context) (func() error, error) {
	rpcSvr := rpc.NewServer()
	api := &AggregatorApi{
		agg: agg,
	}
	if err := rpcSvr.RegisterName("aggregator", api); err != nil {
		return nil, logex.Trace(err)
	}
	rpcSvr.SetBatchLimits(8, 1<<20)
	rpcSvr.SetHTTPBodyLimit(4 << 20)

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", agg.cfg.ListenAddr)
	if err != nil {
		return nil, logex.Trace(err)
	}

	return func() error {
		logex.Infof("listen on: %v", agg.cfg.ListenAddr)
		if err := http.Serve(listener, rpcSvr); err != nil {
			return logex.Trace(err)
		}
		return nil
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	// serveUpdateTask, err := agg.startUpdateOperators(context.Background())
	// if err != nil {
	// 	return logex.Trace(err)
	// }
	// serveUpdateTask()

	serveHttp, err := agg.startRpcServer(ctx)
	if err != nil {
		return logex.Trace(err)
	}

	errChan := make(chan error)
	go func() {
		if err := xmetric.ExportMetricToOpenTelemetry(agg.cfg.OpenTelemetry, agg.Collector); err != nil {
			errChan <- logex.Trace(err)
		}
	}()

	go func() {
		if err := agg.Collector.Serve(agg.cfg.EigenMetricsIpPortAddress); err != nil {
			errChan <- logex.Trace(err)
		}
	}()

	go func() {
		if err := serveHttp(); err != nil {
			errChan <- logex.Trace(err)
		}
	}()

	go func() {
		if err := agg.TaskManager.Run(ctx); err != nil {
			errChan <- logex.Trace(err)
		}
	}()

	for {
		select {
		case response := <-agg.blsAggregationService.GetResponseChannel():
			agg.taskMutex.Lock()
			task := agg.taskIndexMap[response.TaskResponseDigest]
			delete(agg.taskIndexMap, response.TaskResponseDigest)
			agg.taskMutex.Unlock()

			if err := agg.sendAggregatedResponseToContract(ctx, task, response); err != nil {
				logex.Error(err)
			}
		case err := <-errChan:
			logex.Fatal(err)
		}
	}
}

func (agg *Aggregator) submitStateHeader(ctx context.Context, req *TaskRequest) error {
	if req.Task.Identifier.ToInt().Int64() == 1 {
		var md Metadata
		if err := json.Unmarshal([]byte(req.Task.Metadata), &md); err != nil {
			return logex.Trace(err)
		}
		if md.BatchId > 0 {
			if md.BatchId%agg.cfg.Sampling != 0 {
				logex.Infof("[scroll] skip task: %#v", md)
				return nil
			}
		}
	}
	digest, err := req.Task.Digest()
	if err != nil {
		return logex.Trace(err)
	}

	quorumNumbers := make([]types.QuorumNum, len(req.Task.QuorumNumbers))
	quorumThresholdPercentages := make([]types.QuorumThresholdPercentage, len(req.Task.QuorumThresholdPercentages))
	for i, qn := range req.Task.QuorumNumbers {
		quorumNumbers[i] = types.QuorumNum(qn)
		quorumThresholdPercentages[i] = types.QuorumThresholdPercentage(agg.cfg.Threshold)
	}
	req.Task.QuorumThresholdPercentages = types.QuorumThresholdPercentages(quorumThresholdPercentages).UnderlyingType()

	agg.taskMutex.Lock()
	task, ok := agg.taskIndexMap[digest]
	if !ok {
		task = &Task{
			state: req,
			index: agg.taskIndexSeq,
		}
		agg.taskIndexMap[digest] = task
		agg.taskIndexSeq += 1
	}
	agg.taskMutex.Unlock()

	if !ok {
		timeToExpiry := time.Duration(agg.cfg.TimeToExpirySecs) * time.Second
		err = agg.blsAggregationService.InitializeNewTask(task.index, req.Task.ReferenceBlockNumber, quorumNumbers, quorumThresholdPercentages, timeToExpiry)
		if err != nil {
			return logex.Trace(err)
		}
	}

	if err := agg.blsAggregationService.ProcessNewSignature(ctx, task.index, digest, req.Signature, req.OperatorId); err != nil {
		if !strings.Contains(err.Error(), "already completed") {
			return logex.Trace(err)
		}
	}
	return nil
}

func (agg *Aggregator) sendAggregatedResponseToContract(ctx context.Context, task *Task, blsAggServiceResp BlsAggregationServiceResponse) error {
	if blsAggServiceResp.Err != nil {
		return logex.Trace(blsAggServiceResp.Err)
	}

	nonSignerPubkeys := []MultiProverServiceManager.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, bindings.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []MultiProverServiceManager.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, bindings.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := MultiProverServiceManager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        bindings.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        bindings.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	tx, err := agg.multiProverContract.ConfirmState(agg.transactOpt, *task.state.Task.ToAbi(), nonSignerStakesAndSignature)
	if err != nil {
		return logex.Trace(err)
	}
	logex.Pretty(task.state.Task)
	logex.Infof("confirm state: %v", tx.Hash())
	go func() {
		ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				logex.Error(ctx.Err())
				return
			default:
				receipt, _ := agg.client.TransactionReceipt(ctx, tx.Hash())
				if receipt != nil {
					logex.Infof("tx commited: %v, gas used: %v, success: %v", tx.Hash(), receipt.GasUsed, receipt.Status == 1)
					return
				}
				time.Sleep(3 * time.Second)
				continue
			}
		}
	}()
	return nil
}
