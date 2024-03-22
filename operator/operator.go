package operator

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"

	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
)

type Config struct {
	ProverURL     string
	AggregatorURL string

	TaskFetcher *TaskFetcher

	BlsPrivateKey   string
	EcdsaPrivateKey string

	EthRpcUrl string
	EthWsUrl  string

	AVSRegistryCoordinatorAddress common.Address
	OperatorStateRetrieverAddress common.Address
	EigenMetricsIpPortAddress     string
}

type TaskFetcher struct {
	Endpoint  string
	Topics    [][]common.Hash
	Addresses []common.Address
}

type Operator struct {
	cfg      *Config
	nodeApi  *nodeapi.NodeApi
	keyPair  *bls.KeyPair
	ecdsaKey *ecdsa.PrivateKey
	logger   *logex.Logger

	aggregator *aggregator.Client

	operatorId      [32]byte
	operatorAddress common.Address

	proverClient *ProverClient
	eigenClients *clients.Clients
	taskFetcher  *LogTracer
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
	ecdsaPrivateKey, err := crypto.HexToECDSA(cfg.EcdsaPrivateKey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	operatorAddress := crypto.PubkeyToAddress(*ecdsaPrivateKey.Public().(*ecdsa.PublicKey))

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthRpcUrl,
		EthWsUrl:                   cfg.EthWsUrl,
		RegistryCoordinatorAddr:    cfg.AVSRegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: cfg.OperatorStateRetrieverAddress.String(),
		AvsName:                    "multi-prover-operator",
		PromMetricsIpPortAddress:   cfg.EigenMetricsIpPortAddress,
	}

	eigenClients, err := clients.BuildAll(chainioConfig, ecdsaPrivateKey, nil)
	if err != nil {
		return nil, logex.Trace(err)
	}

	operatorId, err := eigenClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	operator := &Operator{
		operatorId:   operatorId,
		keyPair:      kp,
		proverClient: proverClient,
		logger:       logger,
		eigenClients: eigenClients,
		ecdsaKey:     ecdsaPrivateKey,
	}

	client, err := ethclient.Dial(cfg.TaskFetcher.Endpoint)
	if err != nil {
		return nil, logex.Trace(err)
	}
	operator.taskFetcher = NewLogTracer(client, &LogTracerConfig{
		Id:        "operator-log-tracer",
		Wait:      5,
		Max:       100,
		Topics:    cfg.TaskFetcher.Topics,
		Addresses: cfg.TaskFetcher.Addresses,
		Handler:   operator,
	})

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
	blockNumber := new(big.Int).SetBytes(log.Data[:32])
	stateHeader, err := o.proverClient.GetStateHeader(ctx, uint64(blockNumber.Int64()))
	if err != nil {
		return logex.Trace(err)
	}

	// submit to aggregator
	if err := o.aggregator.SubmitStateHeader(ctx, &aggregator.StateHeader{
		StateHeader: stateHeader.StateHeader,
		Signature:   stateHeader.Signature,
		Pubkey:      stateHeader.Pubkey,
		OperatorId:  o.operatorId,
	}); err != nil {
		return logex.Trace(err)
	}

	return nil
}

func (o *Operator) Start(ctx context.Context) error {
	if err := o.RegisterOperatorWithEigenlayer(ctx); err != nil {
		return logex.Trace(err)
	}
	if err := o.RegisterOperatorWithAvs(ctx); err != nil {
		return logex.Trace(err)
	}

	o.logger.Infof("Start Operator... operator info: operatorId=%v, operatorAddr=%v, operatorG1Pubkey=%v, operatorG2Pubkey=%v",
		o.operatorId,
		o.operatorAddress,
		o.keyPair.GetPubKeyG1(),
		o.keyPair.GetPubKeyG2(),
	)

	if err := o.checkIsRegistered(); err != nil {
		return logex.Trace(err)
	}

	if err := o.taskFetcher.Run(ctx); err != nil {
		return logex.Trace(err)
	}

	return nil

}

func (o *Operator) checkIsRegistered() error {
	operatorIsRegistered, err := o.eigenClients.AvsRegistryChainReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	if !operatorIsRegistered {
		return logex.NewErrorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}
	return nil
}

func (o *Operator) RegisterOperatorWithEigenlayer(ctx context.Context) error {
	op := eigenSdkTypes.Operator{
		Address:                 o.operatorAddress.String(),
		EarningsReceiverAddress: o.operatorAddress.String(),
	}
	_, err := o.eigenClients.ElChainWriter.RegisterAsOperator(ctx, op)
	if err != nil {
		o.logger.Errorf("Error registering operator with eigenlayer")
		return err
	}
	return nil
}

func (o *Operator) RegisterOperatorWithAvs(ctx context.Context) error {
	// hardcode these things for now
	quorumNumbers := []eigenSdkTypes.QuorumNum{0}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{123}
	sigValidForSeconds := int64(1_000_000)

	operatorToAvsRegistrationSigExpiry := big.NewInt(time.Now().Unix() + sigValidForSeconds)

	_, err := o.eigenClients.AvsRegistryChainWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		ctx, o.ecdsaKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry, o.keyPair, quorumNumbers, socket,
	)
	if err != nil {
		return logex.Trace(err)
	}
	o.logger.Infof("Registered operator with avs registry coordinator.")

	return nil
}
