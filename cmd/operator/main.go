package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/automata-network/multi-prover-avs/operator"
	"github.com/chzyer/logex"
)

var (
	SemVer    = "0.1.0"
	GitCommit = "(unknown)"
	GitDate   = "(unknown)"
)

type Flag struct {
	Config  string
	Version bool
}

func NewFlag() *Flag {
	var f Flag
	flag.StringVar(&f.Config, "c", "config/operator.json", "config file")
	flag.BoolVar(&f.Version, "v", false, "show version")
	flag.Parse()
	return &f
}

func main() {
	flag := NewFlag()
	if flag.Version {
		fmt.Printf("Version:%v, GitCommit:%v, GitDate:%v\n", SemVer, GitCommit, GitDate)
		return
	}
	o, err := operator.NewOperator(flag.Config)
	if err != nil {
		logex.Fatal(err)
	}
	if err := o.Start(context.Background()); err != nil {
		logex.Fatal(err)
	}
}
