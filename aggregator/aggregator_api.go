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

func (a *AggregatorApi) SubmitTask(ctx context.Context, req *TaskRequest) error {
	// check bls public key
	digest, err := req.Task.Digest()
	if err != nil {
		return logex.Trace(err)
	}

	operatorPubkeys, err := a.agg.registry.GetOperatorsAvsStateAtBlock(ctx, utils.BytesToQuorumNums(req.Task.QuorumNumbers), req.Task.ReferenceBlockNumber)
	if err != nil {
		return logex.Trace(err)
	}
	pubkey := operatorPubkeys[req.OperatorId]

	validPubkey, err := req.Signature.Verify(pubkey.Pubkeys.G2Pubkey, digest)
	if err != nil {
		return logex.Trace(err)
	}
	if !validPubkey {
		return logex.NewErrorf("invalid signature&pubkey")
	}
	x, y := utils.SplitPubkey(pubkey.Pubkeys.G1Pubkey.Serialize())
	pass, err := a.agg.TEELivenessVerifier.VerifyLivenessProof(nil, x, y)
	if err != nil {
		return logex.Trace(err)
	}
	if !pass {
		return logex.NewErrorf("prover not registered")
	}

	if err := a.agg.submitStateHeader(ctx, req); err != nil {
		return logex.Trace(err)
	}
	return nil
}
