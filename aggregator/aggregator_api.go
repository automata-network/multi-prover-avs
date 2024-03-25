package aggregator

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/chzyer/logex"
)

type AggregatorApi struct {
	agg *Aggregator
}

type SignedTaskResponse struct {
	TaskResponse MultiProverServiceManager.IMultiProverServiceManagerStateHeader
	BlsSignature bls.Signature
	OperatorId   types.OperatorId
}

func (a *AggregatorApi) SubmitStateHeader(ctx context.Context, state *StateHeaderRequest) error {
	// check bls public key
	x, y := utils.SplitPubkey(state.Pubkey)
	pass, err := a.agg.sgxVerifier.IsProverRegistered(nil, x, y)
	if err != nil {
		return logex.Trace(err)
	}
	if !pass {
		return logex.NewErrorf("prover not registered")
	}

	if err := a.agg.submitStateHeader(ctx, state); err != nil {
		return logex.Trace(err)
	}
	return nil
}
