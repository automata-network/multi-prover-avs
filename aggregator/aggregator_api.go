package aggregator

import (
	"context"
	"runtime/debug"

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
	defer func() {
		if err := recover(); err != nil {
			logex.Error(err, string(debug.Stack()))
			panic(err)
		}
	}()
	if err := a.submitTask(ctx, req); err != nil {
		logex.Error(err)
		return err
	}
	return nil
}

func (a *AggregatorApi) submitTask(ctx context.Context, req *TaskRequest) error {
	// check bls public key
	digest, err := req.Task.Digest()
	if err != nil {
		return logex.Trace(err)
	}

	operatorPubkeys, err := a.agg.registry.GetOperatorsAvsStateAtBlock(ctx, utils.BytesToQuorumNums(req.Task.QuorumNumbers), req.Task.ReferenceBlockNumber)
	if err != nil {
		return logex.Trace(err)
	}
	pubkey, ok := operatorPubkeys[req.OperatorId]
	if !ok {
		return logex.NewErrorf("operatorId not registered: %v", req.OperatorId)
	}

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
		return logex.NewErrorf("prover[%v] not registered", req.OperatorId)
	}

	if err := a.agg.submitStateHeader(ctx, req); err != nil {
		return logex.Trace(err)
	}
	return nil
}
