package aggregator

import (
	"context"
	"math/big"

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
	logex.Infof("connecting to aggregator: %v", endpoint)
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return &Client{
		client: client,
	}, nil
}

type Metadata struct {
	BatchId    uint64 `json:"batch_id,omitempty"`
	StartBlock uint64 `json:"start_block"`
	EndBlock   uint64 `json:"end_block"`
}

type TaskRequest struct {
	Task       *StateHeader
	Signature  *bls.Signature
	OperatorId types.OperatorId
}

type StateHeader struct {
	Identifier                 *hexutil.Big  `json:"identifier"`
	Metadata                   hexutil.Bytes `json:"metadata"`
	State                      hexutil.Bytes `json:"state"`
	QuorumNumbers              hexutil.Bytes `json:"quorum_numbers"`
	QuorumThresholdPercentages hexutil.Bytes `json:"quorum_threshold_percentages"`
	ReferenceBlockNumber       uint32        `json:"reference_blocknumber"`
}

func (s *StateHeader) ToAbi() *bindings.StateHeader {
	return &bindings.StateHeader{
		CommitteeId:                new(big.Int).Set((*big.Int)(s.Identifier)),
		Metadata:                   []byte(s.Metadata),
		State:                      []byte(s.State),
		QuorumNumbers:              []byte(s.QuorumNumbers),
		QuorumThresholdPercentages: []byte(s.QuorumThresholdPercentages),
		ReferenceBlockNumber:       s.ReferenceBlockNumber,
	}
}

func (s *StateHeader) Digest() (types.TaskResponseDigest, error) {
	digest, err := bindings.DigestStateHeader(s.ToAbi())
	if err != nil {
		return types.TaskResponseDigest{}, logex.Trace(err)
	}

	return digest, nil
}

func (c *Client) SubmitTask(ctx context.Context, task *TaskRequest) error {
	var result bool
	if err := c.client.CallContext(ctx, &result, "aggregator_submitTask", task); err != nil {
		return logex.Trace(err)
	}
	return nil
}
