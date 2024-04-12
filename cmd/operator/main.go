package main

import (
	"context"
	"flag"

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
	o, err := operator.NewOperator(flag.Config)
	if err != nil {
		logex.Fatal(err)
	}
	if err := o.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
