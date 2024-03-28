package aggregator

import (
	"context"
	"math/big"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/TEELivenessVerifier"
	"github.com/automata-network/multi-prover-avs/utils"
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

	EcdsaPrivateKey                    string
	EthHttpEndpoint                    string
	EthWsEndpoint                      string
	MultiProverContractAddress         common.Address
	TEELivenessVerifierContractAddress common.Address

	AVSRegistryCoordinatorAddress common.Address
	OperatorStateRetrieverAddress common.Address
	EigenMetricsIpPortAddress     string
	ScanStartBlock                uint64

	Simulation bool
}

type Aggregator struct {
	cfg *Config

	blsAggregationService blsagg.BlsAggregationService
	transactOpt           *bind.TransactOpts

	client *ethclient.Client

	multiProverContract *MultiProverServiceManager.MultiProverServiceManager
	TEELivenessVerifier *TEELivenessVerifier.TEELivenessVerifierCaller
	registry            *avsregistry.AvsRegistryServiceChainCaller

	taskMutex    sync.Mutex
	taskIndexSeq uint32
	taskIndexMap map[types.TaskResponseDigest]*Task
}

type Task struct {
	state *TaskRequest
	index uint32
}

func NewAggregator(ctx context.Context, cfg *Config) (*Aggregator, error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(cfg.EcdsaPrivateKey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	client, err := ethclient.Dial(cfg.EthHttpEndpoint)
	if err != nil {
		return nil, logex.Trace(err)
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, logex.Trace(err)
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

	operatorPubkeysService, err := NewOperatorPubkeysService(ctx, client, eigenClients.AvsRegistryChainSubscriber, eigenClients.AvsRegistryChainReader, logger, "", cfg.ScanStartBlock, 5000)
	if err != nil {
		return nil, logex.Trace(err)
	}
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(eigenClients.AvsRegistryChainReader, operatorPubkeysService, logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, logger)

	multiProverContract, err := MultiProverServiceManager.NewMultiProverServiceManager(cfg.MultiProverContractAddress, client)
	if err != nil {
		return nil, logex.Trace(err)
	}
	TEELivenessVerifier, err := TEELivenessVerifier.NewTEELivenessVerifierCaller(cfg.TEELivenessVerifierContractAddress, client)
	if err != nil {
		return nil, logex.Trace(err)
	}

	return &Aggregator{
		cfg:                   cfg,
		transactOpt:           transactOpt,
		client:                client,
		blsAggregationService: blsAggregationService,
		multiProverContract:   multiProverContract,
		TEELivenessVerifier:   TEELivenessVerifier,
		registry:              avsRegistryService,
		taskIndexMap:          make(map[types.Bytes32]*Task),
	}, nil
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
	isSimulation, err := agg.TEELivenessVerifier.Simulation(nil)
	if err != nil {
		return logex.Trace(err)
	}
	if isSimulation != agg.cfg.Simulation {
		return logex.NewErrorf("simulation mode not match with the contract: local:%v, remote:%v", agg.cfg.Simulation, isSimulation)
	}

	serveHttp, err := agg.startRpcServer(ctx)
	if err != nil {
		return logex.Trace(err)
	}

	errChan := make(chan error)
	go func() {
		if err := serveHttp(); err != nil {
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

			if err := agg.sendAggregatedResponseToContract(task, response); err != nil {
				logex.Error(err)
			}
		case err := <-errChan:
			logex.Fatal(err)
		}
	}
}

func (agg *Aggregator) submitStateHeader(ctx context.Context, req *TaskRequest) error {
	digest, err := req.Task.Digest()
	if err != nil {
		return logex.Trace(err)
	}
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
		sh := req.Task
		quorumNumbers := make([]types.QuorumNum, len(sh.QuorumNumbers))
		quorumThresholdPercentages := make([]types.QuorumThresholdPercentage, len(sh.QuorumThresholdPercentages))
		for i, qn := range sh.QuorumNumbers {
			quorumNumbers[i] = types.QuorumNum(qn)
		}
		for i, qn := range sh.QuorumThresholdPercentages {
			quorumThresholdPercentages[i] = types.QuorumThresholdPercentage(qn)
		}
		err = agg.blsAggregationService.InitializeNewTask(task.index, sh.ReferenceBlockNumber, quorumNumbers, quorumThresholdPercentages, timeToExpiry)
		if err != nil {
			return logex.Trace(err)
		}
	}

	if err := agg.blsAggregationService.ProcessNewSignature(ctx, task.index, digest, req.Signature, req.OperatorId); err != nil {
		return logex.Trace(err)
	}
	return nil
}

func (agg *Aggregator) sendAggregatedResponseToContract(task *Task, blsAggServiceResp blsagg.BlsAggregationServiceResponse) error {
	if blsAggServiceResp.Err != nil {
		return logex.Trace(blsAggServiceResp.Err)
	}

	nonSignerPubkeys := []MultiProverServiceManager.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []MultiProverServiceManager.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := MultiProverServiceManager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
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
	return nil
}

func ConvertToBN254G1Point(input *bls.G1Point) MultiProverServiceManager.BN254G1Point {
	output := MultiProverServiceManager.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) MultiProverServiceManager.BN254G2Point {
	output := MultiProverServiceManager.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
