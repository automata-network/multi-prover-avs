package operator

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/TEELivenessVerifier"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"

	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
)

type Config struct {
	ProverURL     string
	AggregatorURL string

	TEELivenessVerifierAddr common.Address

	TaskFetcher *TaskFetcher

	BlsPrivateKey   string
	EcdsaPrivateKey string

	EthRpcUrl string
	EthWsUrl  string

	StrategyAddress               common.Address
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
	EigenMetricsIpPortAddress     string
}

type TaskFetcher struct {
	Endpoint  string
	Topics    [][]common.Hash
	Addresses []common.Address
}

type Operator struct {
	cfg        *Config
	blsKeyPair *bls.KeyPair
	ecdsaKey   *ecdsa.PrivateKey
	logger     *logex.Logger

	aggregator *aggregator.Client

	operatorId      [32]byte
	operatorAddress common.Address

	proverClient *ProverClient
	eigenClients *clients.Clients
	taskFetcher  *LogTracer

	ethclient           *ethclient.Client
	TEELivenessVerifier *TEELivenessVerifier.TEELivenessVerifier
}

func NewOperator(cfg *Config) (*Operator, error) {
	kp, err := bls.NewKeyPairFromString(cfg.BlsPrivateKey)
	if err != nil {
		return nil, logex.Trace(err)
	}

	proverClient, err := NewProverClient(cfg.ProverURL)
	if err != nil {
		return nil, logex.Trace(err)
	}

	logger := logex.NewLoggerEx(os.Stderr)
	elog := utils.NewLogger(logger)
	ecdsaPrivateKey, err := crypto.HexToECDSA(cfg.EcdsaPrivateKey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	operatorAddress := crypto.PubkeyToAddress(*ecdsaPrivateKey.Public().(*ecdsa.PublicKey))

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthRpcUrl,
		EthWsUrl:                   cfg.EthWsUrl,
		RegistryCoordinatorAddr:    cfg.RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: cfg.OperatorStateRetrieverAddress.String(),
		AvsName:                    "multi-prover-operator",
		PromMetricsIpPortAddress:   cfg.EigenMetricsIpPortAddress,
	}

	eigenClients, err := clients.BuildAll(chainioConfig, ecdsaPrivateKey, elog)
	if err != nil {
		return nil, logex.Trace(err)
	}

	operatorId, err := eigenClients.AvsRegistryChainReader.GetOperatorId(nil, operatorAddress)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	client, err := ethclient.Dial(cfg.EthRpcUrl)
	if err != nil {
		return nil, logex.Trace(err)
	}
	TEELivenessVerifier, err := TEELivenessVerifier.NewTEELivenessVerifier(cfg.TEELivenessVerifierAddr, client)
	if err != nil {
		return nil, logex.Trace(err)
	}

	operator := &Operator{
		cfg:                 cfg,
		operatorId:          operatorId,
		operatorAddress:     operatorAddress,
		blsKeyPair:          kp,
		proverClient:        proverClient,
		logger:              logger,
		eigenClients:        eigenClients,
		ecdsaKey:            ecdsaPrivateKey,
		TEELivenessVerifier: TEELivenessVerifier,
		ethclient:           client,
	}

	if cfg.TaskFetcher != nil {
		taskFetcherClient, err := ethclient.Dial(cfg.TaskFetcher.Endpoint)
		if err != nil {
			return nil, logex.Trace(err)
		}

		operator.taskFetcher = NewLogTracer(taskFetcherClient, &LogTracerConfig{
			Id:        "operator-log-tracer",
			Wait:      5,
			Max:       100,
			Topics:    cfg.TaskFetcher.Topics,
			Addresses: cfg.TaskFetcher.Addresses,
			Handler:   operator,
		})
	}

	return operator, nil
}

// callback func for task fetcher
func (o *Operator) GetBlock() (uint64, error) {
	return 0, nil
}

// callback func for task fetcher
func (o *Operator) SaveBlock(uint64) error {
	return nil
}

// callback func for task fetcher
func (o *Operator) OnNewLog(ctx context.Context, log *types.Log) error {
	// parse the task
	poe, err := o.proverClient.GetPoe(ctx, log.TxHash)
	if err != nil {
		return logex.Trace(err)
	}
	stateHeader := &aggregator.StateHeader{
		Identifier:                 (*hexutil.Big)(big.NewInt(1)),
		Metadata:                   nil,
		State:                      poe.Pack(),
		QuorumNumbers:              []byte{0},
		QuorumThresholdPercentages: []byte{0},
		ReferenceBlockNumber:       0, // TODO: fixme
	}

	digest, err := stateHeader.Digest()
	if err != nil {
		return logex.Trace(err)
	}
	sig := o.blsKeyPair.SignMessage(digest)

	// submit to aggregator
	if err := o.aggregator.SubmitTask(ctx, &aggregator.TaskRequest{
		Task:       stateHeader,
		Signature:  sig,
		OperatorId: o.operatorId,
	}); err != nil {
		return logex.Trace(err)
	}

	return nil
}

func (o *Operator) Start(ctx context.Context) error {
	if err := o.RegisterOperatorWithEigenlayer(ctx); err != nil {
		return logex.Trace(err)
	}
	if err := o.DepositIntoStrategy(ctx); err != nil {
		return logex.Trace(err)
	}
	if err := o.RegisterOperatorWithAvs(ctx); err != nil {
		return logex.Trace(err)
	}
	if err := o.RegisterAttestationReport(ctx); err != nil {
		return logex.Trace(err)
	}

	o.logger.Infof("Start Operator... operator info: operatorId=%v, operatorAddr=%v, operatorG1Pubkey=%v, operatorG2Pubkey=%v",
		o.operatorId,
		o.operatorAddress,
		o.blsKeyPair.GetPubKeyG1(),
		o.blsKeyPair.GetPubKeyG2(),
	)

	if err := o.checkIsRegistered(); err != nil {
		return logex.Trace(err)
	}

	if o.taskFetcher != nil {
		if err := o.taskFetcher.Run(ctx); err != nil {
			return logex.Trace(err)
		}
	}

	return nil

}

func (o *Operator) checkIsRegistered() error {
	operatorIsRegistered, err := o.eigenClients.AvsRegistryChainReader.IsOperatorRegistered(nil, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	if !operatorIsRegistered {
		return logex.NewErrorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}
	return nil
}

func (o *Operator) RegisterAttestationReport(ctx context.Context) error {
	pubkeyBytes := o.blsKeyPair.PubKey.Serialize()
	if len(pubkeyBytes) != 64 {
		return logex.NewErrorf("invalid pubkey")
	}
	var x, y [32]byte
	copy(x[:], pubkeyBytes[:32])
	copy(y[:], pubkeyBytes[32:64])
	isRegistered, err := o.TEELivenessVerifier.VerifyLivenessProof(nil, x, y)
	if err != nil {
		return logex.Trace(err)
	}
	if isRegistered {
		return nil
	}

	report, err := o.proverClient.GenerateAttestaionReport(ctx, pubkeyBytes)
	if err != nil {
		return logex.Trace(err)
	}
	chainId, err := o.ethclient.ChainID(ctx)
	if err != nil {
		return logex.Trace(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(o.ecdsaKey, chainId)
	if err != nil {
		return logex.Trace(err)
	}

	tx, err := o.TEELivenessVerifier.SubmitLivenessProof(opt, report)
	if err != nil {
		return logex.Trace(err)
	}
	if _, err := utils.WaitTx(ctx, o.ethclient, tx, nil); err != nil {
		return logex.Trace(err)
	}
	return nil
}

func (o *Operator) RegisterOperatorWithEigenlayer(ctx context.Context) error {
	registered, err := o.eigenClients.ElChainReader.IsOperatorRegistered(nil, eigenSdkTypes.Operator{
		Address: o.operatorAddress.String(),
	})
	if err != nil {
		return logex.Trace(err)
	}
	if registered {
		return nil
	}

	op := eigenSdkTypes.Operator{
		Address:                 o.operatorAddress.String(),
		EarningsReceiverAddress: o.operatorAddress.String(),
	}
	receipt, err := o.eigenClients.ElChainWriter.RegisterAsOperator(ctx, op)
	if err != nil {
		return logex.Trace(err, "Error registering operator with eigenlayer")
	}

	o.logger.Infof("Registered operator with Eigenlayer. status: %v", receipt.Status)
	return nil
}

func (o *Operator) DepositIntoStrategy(ctx context.Context) error {
	// strategyAddr common.Address, amount *big.Int
	_, tokenAddr, err := o.eigenClients.ElChainReader.GetStrategyAndUnderlyingToken(nil, o.cfg.StrategyAddress)
	if err != nil {
		return logex.Trace(err, "Failed to fetch strategy contract")
	}
	logex.Info("tokenAddr:", tokenAddr)

	decimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	amount := new(big.Int).Mul(big.NewInt(32), decimal)
	_, err = o.eigenClients.ElChainWriter.DepositERC20IntoStrategy(context.Background(), o.cfg.StrategyAddress, amount)
	if err != nil {
		return logex.Trace(err, "Error depositing into strategy")
	}

	return nil
}

func (o *Operator) RegisterOperatorWithAvs(ctx context.Context) error {
	quorumNumbers := []eigenSdkTypes.QuorumNum{0}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{123}
	sigValidForSeconds := int64(1_000_000)

	operatorToAvsRegistrationSigExpiry := big.NewInt(time.Now().Unix() + sigValidForSeconds)

	_, err := o.eigenClients.AvsRegistryChainWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		ctx, o.ecdsaKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry, o.blsKeyPair, quorumNumbers, socket,
	)
	if err != nil {
		return logex.Trace(err)
	}
	o.logger.Infof("Registered operator with avs registry coordinator.")

	return nil
}
