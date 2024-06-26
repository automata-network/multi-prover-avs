package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"

	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/chzyer/logex"
)

type Flag struct {
	Config string
}

func NewFlag() *Flag {
	var f Flag
	flag.StringVar(&f.Config, "c", "config/aggregator.json", "config file")
	flag.Parse()
	return &f
}

func main() {
	flag := NewFlag()
	cfgBytes, err := os.ReadFile(flag.Config)
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
