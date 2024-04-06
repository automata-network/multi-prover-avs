package operator

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/TEELivenessVerifier"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	Simulation    bool
	Identifier    int64

	TEELivenessVerifierAddr common.Address

	TaskFetcher *TaskFetcher

	BlsPrivateKey   string
	EcdsaPrivateKey string

	EthRpcUrl string
	EthWsUrl  string

	StrategyAddress            common.Address
	RegistryCoordinatorAddress common.Address
	EigenMetricsIpPortAddress  string
}

type TaskFetcher struct {
	Endpoint   string
	Topics     [][]common.Hash
	Addresses  []common.Address
	OffsetFile string
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
	offset       *os.File

	ethclient           *ethclient.Client
	TEELivenessVerifier *TEELivenessVerifier.TEELivenessVerifier
}

func NewOperator(cfg *Config) (*Operator, error) {
	if cfg.Identifier == 0 {
		cfg.Identifier = 1
	}
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
		OperatorStateRetrieverAddr: common.Address{}.String(),
		AvsName:                    "multi-prover-operator",
		PromMetricsIpPortAddress:   cfg.EigenMetricsIpPortAddress,
	}

	eigenClients, err := clients.BuildAll(chainioConfig, ecdsaPrivateKey, elog)
	if err != nil {
		return nil, logex.Trace(err)
	}

	client, err := ethclient.Dial(cfg.EthRpcUrl)
	if err != nil {
		return nil, logex.Trace(err)
	}
	TEELivenessVerifier, err := TEELivenessVerifier.NewTEELivenessVerifier(cfg.TEELivenessVerifierAddr, client)
	if err != nil {
		return nil, logex.Trace(err)
	}

	aggClient, err := aggregator.NewClient(cfg.AggregatorURL)
	if err != nil {
		return nil, logex.Trace(err)
	}

	operator := &Operator{
		cfg:                 cfg,
		operatorAddress:     operatorAddress,
		blsKeyPair:          kp,
		proverClient:        proverClient,
		logger:              logger,
		eigenClients:        eigenClients,
		ecdsaKey:            ecdsaPrivateKey,
		aggregator:          aggClient,
		TEELivenessVerifier: TEELivenessVerifier,
		ethclient:           client,
	}

	if cfg.TaskFetcher != nil {
		taskFetcherClient, err := ethclient.Dial(cfg.TaskFetcher.Endpoint)
		if err != nil {
			return nil, logex.Trace(err)
		}

		offsetFile, err := os.OpenFile(cfg.TaskFetcher.OffsetFile, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, logex.Trace(err)
		}

		operator.offset = offsetFile
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
func (h *Operator) GetBlock() (uint64, error) {
	data := make([]byte, 16)
	n, err := h.offset.ReadAt(data, 0)
	if n == 0 {
		if err == io.EOF {
			return 0, nil
		}
		return 0, logex.Trace(err, n)
	}
	data = bytes.Trim(data[:n], "\x00\r\n ")

	number, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return 0, logex.Trace(err)
	}
	return uint64(number), nil
}

// callback func for task fetcher
func (h *Operator) SaveBlock(offset uint64) error {
	data := []byte(strconv.FormatUint(offset, 10))
	_, err := h.offset.WriteAt(data, 0)
	return err
}

// callback func for task fetcher
func (o *Operator) OnNewLog(ctx context.Context, log *types.Log) error {
	// parse the task
	poe, err := o.proverGetPoe(ctx, log.TxHash, log.Topics)
	if err != nil {
		return logex.Trace(err)
	}

	logex.Pretty(poe)

	blockNumber, err := o.ethclient.BlockNumber(ctx)
	if err != nil {
		return logex.Trace(err)
	}

	stateHeader := &aggregator.StateHeader{
		Identifier:                 (*hexutil.Big)(big.NewInt(o.cfg.Identifier)),
		Metadata:                   nil,
		State:                      poe.Pack(),
		QuorumNumbers:              []byte{0},
		QuorumThresholdPercentages: []byte{0},
		ReferenceBlockNumber:       uint32(blockNumber),
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
	isSimulation, err := o.TEELivenessVerifier.Simulation(nil)
	if err != nil {
		return logex.Trace(err)
	}
	if isSimulation != o.cfg.Simulation {
		return logex.NewErrorf("simulation mode not match with the contract: local:%v, remote:%v", o.cfg.Simulation, isSimulation)
	}

	if err := o.RegisterOperatorWithEigenlayer(ctx); err != nil {
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

	if err := o.taskFetcher.Run(ctx); err != nil {
		return logex.Trace(err)
	}

	return nil
}

func (o *Operator) checkIsRegistered() error {
	operatorIsRegistered, err := o.eigenClients.AvsRegistryChainReader.IsOperatorRegistered(nil, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	if !operatorIsRegistered {
		return logex.NewErrorf("operator is not registered")
	}
	return nil
}

var ABI = func() abi.ABI {
	ty := `[{"inputs":[{"internalType":"uint8","name":"_version","type":"uint8"},{"internalType":"bytes","name":"_parentBatchHeader","type":"bytes"},{"internalType":"bytes[]","name":"_chunks","type":"bytes[]"},{"internalType":"bytes","name":"_skippedL1MessageBitmap","type":"bytes"}],"name":"commitBatch","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	result, err := abi.JSON(bytes.NewReader([]byte(ty)))
	if err != nil {
		panic(err)
	}
	return result
}()

func (o *Operator) proverGetPoe(ctx context.Context, txHash common.Hash, topics []common.Hash) (*Poe, error) {
	if o.cfg.Simulation {

		tx, _, err := o.taskFetcher.source.TransactionByHash(ctx, txHash)
		if err != nil {
			return nil, logex.Trace(err)
		}
		args, err := ABI.Methods["commitBatch"].Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return nil, logex.Trace(err)
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

		startBlockHeader, err := o.taskFetcher.source.HeaderByNumber(ctx, big.NewInt(startBlock))
		if err != nil {
			return nil, logex.Trace(err)
		}
		endBlockHeader, err := o.taskFetcher.source.HeaderByNumber(ctx, big.NewInt(endBlock))
		if err != nil {
			return nil, logex.Trace(err)
		}

		poe := &Poe{
			BatchHash:     topics[2],
			NewStateRoot:  endBlockHeader.Root,
			PrevStateRoot: startBlockHeader.Root,
		}
		return poe, nil
	}
	poe, err := o.proverClient.GetPoe(ctx, txHash)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return poe, nil
}

func (o *Operator) proverGetAttestationReport(ctx context.Context, pubkey []byte) ([]byte, error) {
	if o.cfg.Simulation {
		quote, err := generateSimulationQuote(pubkey)
		if err != nil {
			return nil, logex.Trace(err)
		}
		return quote, nil
	}
	quote, err := o.proverClient.GenerateAttestaionReport(ctx, pubkey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return quote, nil
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

	report, err := o.proverGetAttestationReport(ctx, pubkeyBytes)
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
	logex.Info("registered in TEELivenessVerifier: %v", tx.Hash())
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
	operatorId, err := o.eigenClients.AvsRegistryChainReader.GetOperatorId(nil, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	o.operatorId = operatorId

	if operatorId != [32]byte{} {
		return nil
	}

	if err := o.DepositIntoStrategy(ctx); err != nil {
		return logex.Trace(err)
	}

	quorumNumbers := []eigenSdkTypes.QuorumNum{0}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{}
	if _, err := rand.Read(operatorToAvsRegistrationSigSalt[:]); err != nil {
		return logex.Trace(err)
	}
	sigValidForSeconds := int64(1_000_000)

	operatorToAvsRegistrationSigExpiry := big.NewInt(time.Now().Unix() + sigValidForSeconds)

	if _, err := o.eigenClients.AvsRegistryChainWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		ctx, o.ecdsaKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry, o.blsKeyPair, quorumNumbers, socket,
	); err != nil {
		return logex.Trace(err)
	}
	o.logger.Infof("Registered operator with avs registry coordinator.")

	operatorId, err = o.eigenClients.AvsRegistryChainReader.GetOperatorId(nil, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	o.operatorId = operatorId

	return nil
}
