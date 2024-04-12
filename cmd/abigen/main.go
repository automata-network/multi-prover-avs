package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {
	if err := gen("contracts/bindings", "contracts/out/", "MultiProverServiceManager"); err != nil {
		logex.Fatal(err)
	}
	if err := gen("contracts/bindings", "contracts/dcap-v3-attestation/out/", "AutomataDcapV3Attestation"); err != nil {
		logex.Fatal(err)
	}
	if err := gen("contracts/bindings", "contracts/out/", "TEELivenessVerifier"); err != nil {
		logex.Fatal(err)
	}
	if err := gen("contracts/bindings", "contracts/out/", "ERC20"); err != nil {
		logex.Fatal(err)
	}
}

type Abi struct {
	Abi json.RawMessage
}

func gen(out string, base string, name string) error {
	ty := name
	abiBytes, err := os.ReadFile(filepath.Join(base, name+".sol", name+".json"))
	if err != nil {
		return logex.Trace(err)
	}

	var abiType Abi
	if err := json.Unmarshal(abiBytes, &abiType); err != nil {
		return logex.Trace(err)
	}

	code, err := bind.Bind([]string{ty}, []string{string(abiType.Abi)}, []string{""}, nil, name, bind.LangGo, nil, nil)
	if err != nil {
		return logex.Trace(err)
	}
	outFp := filepath.Join(out, name)
	if err := os.MkdirAll(outFp, 0755); err != nil {
		return logex.Trace(err)
	}
	if err := os.WriteFile(filepath.Join(outFp, name+".go"), []byte(code), 0644); err != nil {
		return logex.Trace(err)
	}
	return nil
}
