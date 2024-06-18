package utils

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const MULTI_PROVER_AVS_NAME = "MultiProverAVS"

func GetAvsName(chainID *big.Int) string {
	name := fmt.Sprintf("%v", chainID)
	switch chainID.Int64() {
	case 1:
		name = "Mainnet"
	case 17000:
		name = "Holesky"
	}
	return fmt.Sprintf("%v_%v", MULTI_PROVER_AVS_NAME, name)
}

func BytesToQuorumNums(data []byte) types.QuorumNums {
	n := make([]types.QuorumNum, len(data))
	for i := range data {
		n[i] = types.QuorumNum(data[i])
	}
	return n
}

func ProverAddrHash(url string) common.Hash {
	url = strings.TrimSpace(url)
	return crypto.Keccak256Hash([]byte(MULTI_PROVER_AVS_NAME), []byte(url))
}

func Retry[T any](retryTime int, sleep time.Duration, fn func() (T, error), ctx ...interface{}) (out T, err error) {
	for i := 0; i < retryTime; i++ {
		if i != 0 {
			time.Sleep(sleep)
		}

		out, err = fn()
		if err != nil {
			logex.Errorf("got error: %v, ctx=%v, retry=%v, wait %v", err, ctx, i, sleep)
			continue
		}
		return out, nil
	}
	return out, logex.Trace(err)
}
