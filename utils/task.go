package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/types"
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
