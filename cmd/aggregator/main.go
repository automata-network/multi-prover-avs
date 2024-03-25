package main

import (
	"context"

	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/chzyer/logex"
)

func main() {
	cfg := &aggregator.Config{
		ListenAddr: ":12345",
	}
	ctx := context.Background()
	agg, err := aggregator.NewAggregator(ctx, cfg)
	if err != nil {
		logex.Fatal(err)
	}
	if err := agg.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
