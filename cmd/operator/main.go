package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"

	"github.com/automata-network/multi-prover-avs/operator"
	"github.com/chzyer/logex"
)

type Flag struct {
	Config string
}

func NewFlag() *Flag {
	var f Flag
	flag.StringVar(&f.Config, "c", "config/operator.json", "config file")
	flag.Parse()
	return &f
}

func main() {
	flag := NewFlag()
	cfgBytes, err := os.ReadFile(flag.Config)
	if err != nil {
		logex.Fatal(err)
	}
	cfg := &operator.Config{}
	if err := json.Unmarshal(cfgBytes, &cfg); err != nil {
		logex.Fatal(err)
	}
	o, err := operator.NewOperator(cfg)
	if err != nil {
		logex.Fatal(err)
	}
	if err := o.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
