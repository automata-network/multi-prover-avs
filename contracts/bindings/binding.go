package bindings

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/BLSApkRegistry"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/ERC20"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/TEELivenessVerifier"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

var MultiProverABI = func() *abi.ABI {
	abi, err := MultiProverServiceManager.MultiProverServiceManagerMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	return abi
}()

var TEELivenessVerifierABI = func() *abi.ABI {
	abi, err := TEELivenessVerifier.TEELivenessVerifierMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	return abi
}()

var ERC20ABI = func() *abi.ABI {
	abi, err := ERC20.ERC20MetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	return abi
}()

type Binding struct {
}

type StateHeader = MultiProverServiceManager.IMultiProverServiceManagerStateHeader
type ReducedStateHeader = MultiProverServiceManager.IMultiProverServiceManagerReducedStateHeader

type BlsApkRegistryGetter interface {
	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)
}

func GetBlsApkRegistryCaller(caller bind.ContractCaller, getter BlsApkRegistryGetter) (*BLSApkRegistry.BLSApkRegistryCaller, error) {
	blsApkRegistryAddr, err := getter.BlsApkRegistry(nil)
	if err != nil {
		return nil, logex.Trace(err)
	}
	blsApkRegistry, err := BLSApkRegistry.NewBLSApkRegistryCaller(blsApkRegistryAddr, caller)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return blsApkRegistry, nil
}

func GetOperatorAddrFromBlsKey(blskey *bls.KeyPair, caller bind.ContractCaller, getter BlsApkRegistryGetter) (common.Address, error) {
	blsApkRegistry, err := GetBlsApkRegistryCaller(caller, getter)
	if err != nil {
		return common.Address{}, logex.Trace(err)
	}
	operatorId := eigenSdkTypes.OperatorIdFromKeyPair(blskey)
	blsBindOperatorAddr, err := blsApkRegistry.GetOperatorFromPubkeyHash(nil, operatorId)
	if err != nil {
		return common.Address{}, logex.Trace(err)
	}
	return blsBindOperatorAddr, nil
}

func GetOperatorAddrByOperatorID(caller bind.ContractCaller, getter BlsApkRegistryGetter, operatorId [32]byte) (common.Address, error) {
	blsApkRegistry, err := GetBlsApkRegistryCaller(caller, getter)
	if err != nil {
		return common.Address{}, logex.Trace(err)
	}
	operatorAddr, err := blsApkRegistry.GetOperatorFromPubkeyHash(nil, operatorId)
	if err != nil {
		return common.Address{}, logex.Trace(err)
	}
	return operatorAddr, nil
}

func DigestStateHeader(s *StateHeader) (types.TaskResponseDigest, error) {
	reduced := ReducedStateHeader{s.CommitteeId, s.Metadata, s.State, s.ReferenceBlockNumber}
	argTypes := MultiProverABI.Methods["_hashReducedStateHeader"].Inputs[:1]
	digest, err := argTypes.Pack(reduced)
	if err != nil {
		return types.TaskResponseDigest{}, logex.Trace(err)
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(digest)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

func Keccak256(data []byte) common.Hash {
	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])
	return taskResponseDigest
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

type JsonError interface {
	Error() string
	ErrorCode() int
	ErrorData() interface{}
}

func MultiProverError(err error) error {
	return DecodeError(MultiProverABI, err)
}

func DecodeError(abi *abi.ABI, err error) error {
	je, ok := err.(JsonError)
	if !ok {
		return err
	}
	errorData, ok := je.ErrorData().(string)
	if !ok {
		return err
	}
	data, er := hex.DecodeString(strings.TrimPrefix(errorData, "0x"))
	if er == nil {
		for name, er := range abi.Errors {
			if bytes.Equal(er.ID[:4], data) {
				return logex.NewErrorf("%v: %v", je.Error(), name)
			}
		}
	}
	return logex.NewErrorf("%v: %v", je.Error(), errorData)
}

func ReportDataBytes(reportData *TEELivenessVerifier.TEELivenessVerifierReportDataV2) ([]byte, error) {
	method, ok := TEELivenessVerifierABI.Methods["submitLivenessProofV2"]
	if !ok {
		return nil, logex.NewErrorf("method submitLivenessProofV2 not found in ABI")
	}
	args := abi.Arguments(method.Inputs[:1])
	data, err := args.Pack(reportData)
	if err != nil {
		return nil, logex.Trace(err, "PackReportDataV2")
	}
	return data, nil
}

func ReportDataHash(reportData *TEELivenessVerifier.TEELivenessVerifierReportDataV2) (common.Hash, error) {
	data, err := ReportDataBytes(reportData)
	if err != nil {
		return common.Hash{}, logex.Trace(err)
	}
	return Keccak256(data), nil
}
