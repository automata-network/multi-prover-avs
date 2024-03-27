package utils

import "github.com/Layr-Labs/eigensdk-go/types"

func BytesToQuorumNums(data []byte) types.QuorumNums {
	n := make([]types.QuorumNum, len(data))
	for i := range data {
		n[i] = types.QuorumNum(data[i])
	}
	return n
}
