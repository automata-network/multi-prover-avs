package operator

import (
	"context"
	"encoding/json"

	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type ProverClient struct {
	client *rpc.Client
}

func NewProverClient(url string) (*ProverClient, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return &ProverClient{
		client: client,
	}, nil
}

type Poe struct {
	BatchHash      common.Hash
	StateHash      common.Hash
	PrevStateRoot  common.Hash
	NewStateRoot   common.Hash
	WithdrawalRoot common.Hash
}

func (p *Poe) Pack() []byte {
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ProverClient) GenerateAttestaionReport(ctx context.Context, pubkey hexutil.Bytes) (hexutil.Bytes, error) {
	if len(pubkey) != 64 {
		return nil, logex.NewErrorf("invalid pubkey")
	}
	var attestationReport hexutil.Bytes
	if err := p.client.CallContext(ctx, &attestationReport, "generateAttestationReport", pubkey); err != nil {
		return nil, logex.Trace(err)
	}
	return attestationReport, nil
}

func (p *ProverClient) GetPoe(ctx context.Context, blockNumber uint64) (*Poe, error) {
	var result *Poe
	if err := p.client.CallContext(ctx, &result, "getPoe", blockNumber); err != nil {
		return nil, logex.Trace(err)
	}
	return result, nil
}
