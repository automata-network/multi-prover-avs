package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/automata-network/multi-prover-avs/operator"
	"github.com/chzyer/logex"
)

func main() {
	cfgBytes, err := os.ReadFile("config/operator.json")
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
