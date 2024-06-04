package utils

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WeiToF64(val *big.Int, decimals int64) float64 {
	decimalF := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	amountF := new(big.Float).SetInt(val)
	amountF = amountF.Quo(amountF, decimalF)
	out, _ := amountF.Float64()
	return out
}

func GetTxCost(receipt *types.Receipt) float64 {
	balanceCost := new(big.Int).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))
	return WeiToF64(balanceCost, 18)
}

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
