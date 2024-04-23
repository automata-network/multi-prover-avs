package operator

import (
	"math/big"
	"testing"

	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
)

func TestAbi(t *testing.T) {
	abi, err := MultiProverServiceManager.MultiProverServiceManagerMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	method := abi.Methods["confirmState"]
	confirmState := method.Inputs[:1]
	stateHeader := MultiProverServiceManager.IMultiProverServiceManagerStateHeader{
		CommitteeId: big.NewInt(1),
	}
	_ = stateHeader
	var test MultiProverServiceManager.BN254G1Point
	data, err := confirmState.Pack(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
