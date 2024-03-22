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
	agg := aggregator.NewAggregator(cfg)
	if err := agg.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
