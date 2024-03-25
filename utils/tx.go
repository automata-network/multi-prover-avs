package utils

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitTx(ctx context.Context, client *ethclient.Client, tx *types.Transaction, deferFunc func()) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer func() {
		if deferFunc != nil {
			deferFunc()
		}
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			receipt, _ := client.TransactionReceipt(ctx, tx.Hash())
			if receipt != nil {
				return receipt, nil
			}
			time.Sleep(5 * time.Second)
			continue
		}
	}
}
