package bindings

import (
	"github.com/Layr-Labs/eigensdk-go/types"
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

type Binding struct {
}

type StateHeader = MultiProverServiceManager.IMultiProverServiceManagerStateHeader
type ReducedStateHeader = MultiProverServiceManager.IMultiProverServiceManagerReducedStateHeader

func DigestStateHeader(s *StateHeader) (types.TaskResponseDigest, error) {
	reduced := ReducedStateHeader{s.Identifier, s.Metadata, s.State, s.ReferenceBlockNumber}
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
