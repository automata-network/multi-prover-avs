package main

import (
	"context"

	"github.com/automata-network/multi-prover-avs/operator"
	"github.com/chzyer/logex"
)

func main() {
	cfg := &operator.Config{}
	o, err := operator.NewOperator(cfg)
	if err != nil {
		logex.Fatal(err)
	}
	if err := o.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
