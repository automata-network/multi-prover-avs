package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/ERC20"
	"github.com/automata-network/multi-prover-avs/operator"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/chzyer/flagly"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
)

var (
	SemVer    = "0.1.0"
	GitCommit = "(unknown)"
	GitDate   = "(unknown)"
)

type OprTool struct {
	OptIn   *OprToolOptIn   `flagly:"handler"`
	OptOut  *OprToolOptOut  `flagly:"handler"`
	Deposit *OprToolDeposit `flagly:"handler"`
}

type OprToolOptIn struct {
	Config       string `default:"config/operator.json"`
	EcdsaKeyPath string `default:"~/.eigenlayer/operator_keys/operator.ecdsa.key.json"`
	Socket       string `default:"Not Needed"`
	Quorums      string `default:"0"`
	SigValidSecs int64  `default:"1000000"`
	Version      bool   `name:"v"`
}

func parseQuorums(n string) ([]eigenSdkTypes.QuorumNum, error) {
	var out []eigenSdkTypes.QuorumNum
	sp := strings.Split(n, ",")
	for _, item := range sp {
		val, err := strconv.Atoi(item)
		if err != nil {
			return nil, logex.Trace(err, item)
		}
		out = append(out, eigenSdkTypes.QuorumNum(val))
	}
	if len(out) == 0 {
		return nil, logex.NewErrorf("empty quorum")
	}
	return out, nil
}

func (o *OprToolOptIn) FlaglyHandle() error {
	if o.Version {
		fmt.Printf("Version:%v, GitCommit:%v, GitDate:%v\n", SemVer, GitCommit, GitDate)
		return nil
	}
	ecdsaKey, err := utils.PromptEcdsaKey(o.EcdsaKeyPath)
	if err != nil {
		return logex.Trace(err)
	}
	ctx, err := operator.ParseConfigContext(o.Config, ecdsaKey)
	if err != nil {
		return logex.Trace(err)
	}

	quorumNumbers, err := parseQuorums(o.Quorums)
	if err != nil {
		return logex.Trace(err)
	}
	operatorToAvsRegistrationSigSalt := [32]byte{}
	if _, err := rand.Read(operatorToAvsRegistrationSigSalt[:]); err != nil {
		return logex.Trace(err)
	}
	sigValidForSeconds := int64(o.SigValidSecs)

	operatorToAvsRegistrationSigExpiry := big.NewInt(time.Now().Unix() + sigValidForSeconds)

	receipt, err := ctx.EigenClients.AvsRegistryChainWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		context.Background(), ecdsaKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry, ctx.BlsKey, quorumNumbers, o.Socket,
	)
	if err != nil {
		return logex.Trace(bindings.MultiProverError(err))
	}
	logex.Infof("Registered operator with avs registry coordinator, succ: %v", receipt.Status == 1)

	operatorId, err := ctx.EigenClients.AvsRegistryChainReader.GetOperatorId(nil, utils.EcdsaAddress(ecdsaKey))
	if err != nil {
		return logex.Trace(err)
	}

	logex.Infof("operatorID: %v", hex.EncodeToString(operatorId[:]))
	return nil
}

type OprToolOptOut struct {
	EcdsaKeyPath string `default:"~/.eigenlayer/operator_keys/operator.ecdsa.key.json"`
	Config       string `default:"config/operator.json"`
	Quorums      string `default:"0"`
	Version      bool   `name:"v"`
}

func (o *OprToolOptOut) FlaglyHandle() error {
	if o.Version {
		fmt.Printf("Version:%v, GitCommit:%v, GitDate:%v\n", SemVer, GitCommit, GitDate)
		return nil
	}
	ecdsaKey, err := utils.PromptEcdsaKey(o.EcdsaKeyPath)
	if err != nil {
		return logex.Trace(err)
	}
	ctx, err := operator.ParseConfigContext(o.Config, ecdsaKey)
	if err != nil {
		return logex.Trace(err)
	}

	quorumNumbers, err := parseQuorums(o.Quorums)
	if err != nil {
		return logex.Trace(err)
	}

	receipt, err := ctx.EigenClients.AvsRegistryChainWriter.DeregisterOperator(
		context.Background(),
		quorumNumbers,
		regcoord.BN254G1Point(bindings.ConvertToBN254G1Point(ctx.BlsKey.GetPubKeyG1())),
	)
	if err != nil {
		return logex.Trace(bindings.MultiProverError(err))
	}

	logex.Infof("tx: %v, succ: %v", receipt.TxHash, receipt.Status == 1)
	return nil
}

type OprToolDeposit struct {
	EcdsaKeyPath    string `default:"~/.eigenlayer/operator_keys/operator.ecdsa.key.json"`
	Config          string `default:"config/operator.json"`
	StrategyAddress string `name:"strategy"`
	Amount          string `default:"32"`
	Check           bool
	Version         bool `name:"v"`
}

func (o *OprToolDeposit) FlaglyHandle() error {
	if o.Version {
		fmt.Printf("Version:%v, GitCommit:%v, GitDate:%v\n", SemVer, GitCommit, GitDate)
		return nil
	}
	ecdsaKey, err := utils.PromptEcdsaKey(o.EcdsaKeyPath)
	if err != nil {
		return logex.Trace(err)
	}
	operatorAddress := utils.EcdsaAddress(ecdsaKey)
	ctx, err := operator.ParseConfigContext(o.Config, ecdsaKey)
	if err != nil {
		return logex.Trace(err)
	}
	strategyAddress := common.HexToAddress(o.StrategyAddress)
	var empty common.Address
	if strategyAddress == empty {
		return flagly.ErrShowUsage
	}
	registered, err := ctx.EigenClients.ElChainReader.IsOperatorRegistered(nil, eigenSdkTypes.Operator{
		Address: operatorAddress.String(),
	})
	if err != nil {
		return logex.Trace(bindings.MultiProverError(err))
	}
	if !registered {
		return logex.NewErrorf("operator[%v] not registered", operatorAddress)
	}

	_, tokenAddr, err := ctx.EigenClients.ElChainReader.GetStrategyAndUnderlyingToken(nil, strategyAddress)
	if err != nil {
		err = bindings.MultiProverError(err)
		return logex.Trace(err, "Failed to fetch strategy contract")
	}

	erc20Caller, err := ERC20.NewERC20Caller(tokenAddr, ctx.Client)
	if err != nil {
		return logex.Trace(err)
	}
	decimal, err := erc20Caller.Decimals(nil)
	if err != nil {
		return logex.Trace(err)
	}
	decimalF := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil))
	amountF, ok := new(big.Float).SetString(o.Amount)
	if !ok {
		return logex.NewError("parseAmount")
	}

	amountF = amountF.Mul(amountF, decimalF)

	amount, _ := amountF.Int(nil)
	logex.Infof("deposit token address: %v, decimals: %v, amount: %v", tokenAddr, decimal, amount)

	if !o.Check {
		receipt, err := ctx.EigenClients.ElChainWriter.DepositERC20IntoStrategy(context.Background(), strategyAddress, amount)
		if err != nil {
			return logex.Trace(bindings.MultiProverError(err), "Error depositing into strategy")
		}
		logex.Infof("tx: %v, succ: %v", receipt.TxHash, receipt.Status == 1)
	}

	shares, err := ctx.EigenClients.ElChainReader.GetOperatorSharesInStrategy(nil, operatorAddress, strategyAddress)
	if err != nil {
		return logex.Trace(bindings.MultiProverError(err))
	}
	logex.Infof("current shares: %v", shares)
	return nil
}

func main() {
	if err := flagly.RunByArgs(&OprTool{}, os.Args); err != nil {
		logex.Fatal(err)
	}
}
