package bindings

import (
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

func PackStateHeader(s *StateHeader) ([]byte, error) {
	argTypes := MultiProverABI.Methods["confirmState"].Inputs[:1]
	return argTypes.Pack(s)
}
