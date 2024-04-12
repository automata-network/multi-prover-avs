package bindings

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/ERC20"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

var MultiProverABI = func() *abi.ABI {
	abi, err := MultiProverServiceManager.MultiProverServiceManagerMetaData.GetAbi()
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
