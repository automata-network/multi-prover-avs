package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/chzyer/logex"
)

func main() {
	cfgBytes, err := os.ReadFile("config/aggregator.json")
	if err != nil {
		logex.Fatal(err)
	}

	var cfg aggregator.Config
	if err := json.Unmarshal(cfgBytes, &cfg); err != nil {
		logex.Fatal(err)
	}
	ctx := context.Background()
	agg, err := aggregator.NewAggregator(ctx, &cfg)
	if err != nil {
		logex.Fatal(err)
	}
	if err := agg.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
