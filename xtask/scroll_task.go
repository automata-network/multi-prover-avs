package xtask

import (
	"bytes"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var ScrollABI = func() abi.ABI {
	ty := `[{"inputs":[{"internalType":"uint8","name":"_version","type":"uint8"},{"internalType":"bytes","name":"_parentBatchHeader","type":"bytes"},{"internalType":"bytes[]","name":"_chunks","type":"bytes[]"},{"internalType":"bytes","name":"_skippedL1MessageBitmap","type":"bytes"}],"name":"commitBatch","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	result, err := abi.JSON(bytes.NewReader([]byte(ty)))
	if err != nil {
		panic(err)
	}
	return result
}()

type ScrollContext struct {
	Hash      common.Hash     `json:"hash"`
	Interning json.RawMessage `json:"interning"`
	Pob       json.RawMessage `json:"pob"`
}

type ScrollTaskExt struct {
	StartBlock           *hexutil.Big  `json:"start_block"`
	EndBlock             *hexutil.Big  `json:"end_block"`
	BatchData            hexutil.Bytes `json:"batch_data"`
	CommitTx             common.Hash   `json:"commit_tx"`
	ReferenceBlockNumber uint64        `json:"reference_block_number"`
}

type LineaTaskExt struct {
	StartBlock              *hexutil.Big `json:"start_block"`
	EndBlock                *hexutil.Big `json:"end_block"`
	CommitTx                common.Hash  `json:"commit_tx"`
	PrevCommitTx            common.Hash  `json:"prev_commit_tx"`
	PrevBatchFinalStateRoot common.Hash  `json:"prev_batch_final_state_root"`
	FinalStateRoot          common.Hash  `json:"final_state_root"`
	ReferenceBlockNumber    uint64       `json:"reference_block_number"`
}

type LineaContext struct {
	Hash      common.Hash     `json:"hash"`
	Interning json.RawMessage `json:"interning"`
	Pob       json.RawMessage `json:"pob"`
}
