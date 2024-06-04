package operator

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"os"
	"strings"

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

type ConfigContext struct {
	Config              *Config
	BlsKey              *bls.KeyPair
	AttestationEcdsaKey *ecdsa.PrivateKey
	Client              *ethclient.Client
	AttestationClient   *ethclient.Client
	EigenClients        *clients.Clients
	AvsName             string
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
			return nil, logex.Trace(err, "generate test ecdsa key", cfgPath)
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
	cfg.InitFromEnv()

	kp, err := utils.ReadBlsKey(cfg.BlsKeyFile, cfg.BlsKeyPassword)
	if err != nil {
		return nil, logex.Trace(err, "BlsKeyFile", cfg.BlsKeyFile, cfgPath)
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

	avsChainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, logex.Trace(err, "check chain ID")
	}

	avsName := utils.GetAvsName(avsChainID)

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthRpcUrl,
		EthWsUrl:                   cfg.EthRpcUrl,
		RegistryCoordinatorAddr:    cfg.RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: common.Address{}.String(),
		AvsName:                    avsName,
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
		AvsName:             avsName,
	}, nil
}

type Config struct {
	ProverURL     string
	AggregatorURL string
	Identifier    int64

	BlsKeyFile     string
	BlsKeyPassword string

	EthRpcUrl string

	AttestationLayerEcdsaKey string
	AttestationLayerRpcURL   string

	RegistryCoordinatorAddress common.Address
	TEELivenessVerifierAddress common.Address
	EigenMetricsIpPortAddress  string
	NodeApiIpPortAddress       string
}

func (c *Config) InitFromEnv() {
	if c.NodeApiIpPortAddress == "" {
		c.NodeApiIpPortAddress = ":15692"
	}
}

func (c *Config) check(env *string) {
	if strings.HasPrefix(*env, "$") {
		envVal := os.Getenv((*env)[1:])
		if envVal != "" {
			*env = envVal
		}
	}
}

type TaskFetcher struct {
	Endpoint         string
	Topics           [][]common.Hash
	Addresses        []common.Address
	OffsetFile       string
	ScanIntervalSecs int64
}
