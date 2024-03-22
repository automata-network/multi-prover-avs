package aggregator

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/contracts/bindings"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	client *rpc.Client
}

func NewClient(endpoint string) (*Client, error) {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return &Client{
		client: client,
	}, nil
}

type StateHeader struct {
	StateHeader *bindings.StateHeader
	Signature   bls.Signature
	Pubkey      hexutil.Bytes
	OperatorId  types.OperatorId
}

func (s *StateHeader) Digest() (types.TaskResponseDigest, error) {
	digest, err := bindings.PackStateHeader(s.StateHeader)
	if err != nil {
		return types.TaskResponseDigest{}, logex.Trace(err)
	}
	return types.TaskResponseDigest(digest), nil
}

func (c *Client) SubmitStateHeader(ctx context.Context, state *StateHeader) error {
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

	var result bool
	if err := c.client.CallContext(ctx, &result, "aggregator_submitStateHeader", state); err != nil {
		return logex.Trace(err)
	}
	return nil
}
