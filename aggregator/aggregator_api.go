package aggregator

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/MultiProverServiceManager"
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

func (a *AggregatorApi) SubmitStateHeader(ctx context.Context, state *StateHeader) error {
	// check bls public key
	pubkey := (&bls.G2Point{}).Deserialize(state.Pubkey)
	digest, err := state.Digest()
	if err != nil {
		return logex.Trace(err)
	}
	pass, err := state.Signature.Verify(pubkey, digest)
	if err != nil {
		return logex.Trace(err)
	}
	if !pass {
		return logex.NewErrorf("signature validation failed")
	}
	pass, err = a.agg.sgxVerifier.IsProverRegistered(nil, pubkey.X.A0.Bytes(), pubkey.Y.A0.Bytes())
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
