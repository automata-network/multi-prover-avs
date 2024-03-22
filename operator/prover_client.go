package operator

import (
	"context"
	"encoding/json"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/automata-network/multi-prover-avs/contracts/bindings"
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

	// signature
}

func (p *Poe) Pack() []byte {
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return data
}

type StateHeader struct {
	StateHeader *bindings.StateHeader
	Signature   bls.Signature
	Pubkey      hexutil.Bytes
}

func (p *ProverClient) GetStateHeader(ctx context.Context, blockNumber uint64) (*StateHeader, error) {
	var result *StateHeader
	if err := p.client.CallContext(ctx, &result, "getStateHeader", blockNumber); err != nil {
		return nil, logex.Trace(err)
	}
	return result, nil
}
