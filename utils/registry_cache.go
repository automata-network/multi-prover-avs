package utils

import (
	"context"
	"fmt"
	"time"

	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var _ avsregistry.AvsRegistryService = (*RegistryCache)(nil)

type OperatorAvsStateMap = map[types.OperatorId]types.OperatorAvsState
type QuorumAvsStateMap = map[types.QuorumNum]types.QuorumAvsState

type RegistryCache struct {
	avsStateCaller *avsregistry.AvsRegistryServiceChainCaller

	operatorAvsStateCache *SingleFlightLruCache[string, OperatorAvsStateMap]
	quorumAvsStateCache   *SingleFlightLruCache[string, QuorumAvsStateMap]
}

func NewRegistryCache(avsStateCaller *avsregistry.AvsRegistryServiceChainCaller) *RegistryCache {
	operatorAvsStateCache := NewSingleFlightLruCache[string, OperatorAvsStateMap](10)
	quorumAvsStateCache := NewSingleFlightLruCache[string, QuorumAvsStateMap](10)
	return &RegistryCache{
		avsStateCaller: avsStateCaller,

		operatorAvsStateCache: operatorAvsStateCache,
		quorumAvsStateCache:   quorumAvsStateCache,
	}
}

func (c *RegistryCache) GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers types.QuorumNums, blockNumber types.BlockNum) (QuorumAvsStateMap, error) {
	key := fmt.Sprintf("%v_%v", quorumNumbers, blockNumber)
	avsState, err := c.quorumAvsStateCache.Get(key, func(key string) (QuorumAvsStateMap, error) {
		return Retry(5, time.Second, func() (QuorumAvsStateMap, error) {
			states, err := c.avsStateCaller.GetQuorumsAvsStateAtBlock(ctx, quorumNumbers, blockNumber)
			if err != nil {
				return nil, logex.Trace(err)
			}
			return states, nil
		})
	})
	if err != nil {
		return nil, logex.Trace(err, key)
	}

	return avsState, nil
}

func (c *RegistryCache) GetCheckSignaturesIndices(opts *bind.CallOpts, referenceBlockNumber types.BlockNum, quorumNumbers types.QuorumNums, nonSignerOperatorIds []types.OperatorId) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	return c.avsStateCaller.GetCheckSignaturesIndices(opts, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

func (c *RegistryCache) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers types.QuorumNums, blockNumber types.BlockNum) (OperatorAvsStateMap, error) {
	key := fmt.Sprintf("%v_%v", quorumNumbers, blockNumber)
	avsState, err := c.operatorAvsStateCache.Get(key, func(key string) (OperatorAvsStateMap, error) {
		return Retry(5, time.Second, func() (OperatorAvsStateMap, error) {
			states, err := c.avsStateCaller.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, blockNumber)
			if err != nil {
				return nil, logex.Trace(err)
			}
			return states, nil
		})
	})
	if err != nil {
		return nil, logex.Trace(err, key)
	}

	return avsState, nil
}
