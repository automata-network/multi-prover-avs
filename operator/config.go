package operator

import (
	"crypto/ecdsa"
	"encoding/json"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/automata-network/multi-prover-avs/contracts/bindings"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/RegistryCoordinator"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	AppName   = "multi-prover-operator"
	SemVer    = "0.1.0"
	GitCommit = ""
	GitDate   = ""
)

type ConfigContext struct {
	Config              *Config
	BlsKey              *bls.KeyPair
	AttestationEcdsaKey *ecdsa.PrivateKey
	Client              *ethclient.Client
	AttestationClient   *ethclient.Client
	EigenClients        *clients.Clients
}

func (c *ConfigContext) QueryOperatorAddress() (common.Address, error) {
	registry, err := RegistryCoordinator.NewRegistryCoordinatorCaller(c.Config.RegistryCoordinatorAddress, c.Client)
	if err != nil {
		return utils.ZeroAddress, logex.Trace(err)
	}
	operatorAddress, err := bindings.GetOperatorAddrFromBlsKey(c.BlsKey, c.Client, registry)
	if err != nil {
		return utils.ZeroAddress, logex.Trace(err)
	}
	return operatorAddress, nil
}

func ParseConfigContext(cfgPath string, ecdsaKey *ecdsa.PrivateKey) (*ConfigContext, error) {
	if ecdsaKey == nil {
		var err error
		ecdsaKey, err = crypto.GenerateKey()
		if err != nil {
			return nil, logex.Trace(err, "generate test ecdsa key")
		}
	}
	var cfg Config
	cfgData, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, logex.Trace(err)
	}
	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		return nil, logex.Trace(err, cfgPath)
	}

	kp, err := utils.ReadBlsKey(cfg.BlsKeyFile, cfg.BlsKeyPassword)
	if err != nil {
		return nil, logex.Trace(err, cfg.BlsKeyFile)
	}

	logger := logex.NewLoggerEx(os.Stderr)
	elog := utils.NewLogger(logger)

	logex.Debugf("connecting to EthRpcUrl %v...", cfg.EthRpcUrl)
	client, err := ethclient.Dial(cfg.EthRpcUrl)
	if err != nil {
		return nil, logex.Trace(err, cfg.EthRpcUrl)
	}
	logex.Debugf("connecting to AttestationLayerRpcURL %v...", cfg.AttestationLayerRpcURL)
	attestationClient, err := ethclient.Dial(cfg.AttestationLayerRpcURL)
	if err != nil {
		return nil, logex.Trace(err, "AttestationLayerRpcURL", cfg.AttestationLayerRpcURL)
	}

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthRpcUrl,
		EthWsUrl:                   cfg.EthWsUrl,
		RegistryCoordinatorAddr:    cfg.RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: common.Address{}.String(),
		AvsName:                    AppName,
		PromMetricsIpPortAddress:   cfg.EigenMetricsIpPortAddress,
	}

	logex.Debugf("build clients %#v", chainioConfig)
	eigenClients, err := clients.BuildAll(chainioConfig, ecdsaKey, elog)
	if err != nil {
		return nil, logex.Trace(err)
	}

	if cfg.Identifier == 0 {
		cfg.Identifier = 1
	}

	attestationEcdsaKey, err := crypto.HexToECDSA(cfg.AttestationLayerEcdsaKey)
	if err != nil {
		return nil, logex.Trace(err, "AttestationLayerEcdsaKey")
	}

	return &ConfigContext{
		Config:              &cfg,
		BlsKey:              kp,
		Client:              client,
		EigenClients:        eigenClients,
		AttestationEcdsaKey: attestationEcdsaKey,
		AttestationClient:   attestationClient,
	}, nil
}

type Config struct {
	ProverURL     string
	AggregatorURL string
	Simulation    bool
	Identifier    int64

	TaskFetcher *TaskFetcher

	BlsKeyFile     string
	BlsKeyPassword string

	EthRpcUrl string
	EthWsUrl  string

	AttestationLayerEcdsaKey string
	AttestationLayerRpcURL   string

	StrategyAddress            common.Address
	RegistryCoordinatorAddress common.Address
	TEELivenessVerifierAddress common.Address
	EigenMetricsIpPortAddress  string
}

type TaskFetcher struct {
	Endpoint         string
	Topics           [][]common.Hash
	Addresses        []common.Address
	OffsetFile       string
	ScanIntervalSecs int64
}
