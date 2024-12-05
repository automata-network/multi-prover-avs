package utils

import (
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
)

const ATT_AUTOMATA = "automata"
const ATT_AUTOMATA_TEST = "automata_test"
const ATT_OPTIMISM = "optimism"
const ATT_HOLESKY = "holesky"

type PresetConfig struct {
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
	// TEELivenessVerifierAddress    common.Address
	DefaultAttestationLayer string
	AttestationLayer        map[string]*PresetAttestationConfig
	LineaProverURL          string
	ScrollProverURL         string
}

func (p *PresetConfig) GetAttestationLayer(profile string) *PresetAttestationConfig {
	if profile == "" {
		profile = p.DefaultAttestationLayer
	}
	cfg, ok := p.AttestationLayer[profile]
	if !ok {
		logex.Fatal("unknown Attestation Layer Profile: %q", profile)
	}
	return cfg
}

type PresetAttestationConfig struct {
	Address common.Address
	URL     string
}

var PresetConfigs = []*PresetConfig{HoleskyTestnetPreset, MainnetPreset}

var HoleskyTestnetPreset = &PresetConfig{
	RegistryCoordinatorAddress:    common.HexToAddress("0x62c715575cE3Ad7C5a43aA325b881c70564f2215"),
	OperatorStateRetrieverAddress: common.HexToAddress("0xbfd43ac0a19c843e44491c3207ea13914818E214"),
	LineaProverURL:                "https://avs-prover-staging.ata.network",
	ScrollProverURL:               "https://avs-prover-staging.ata.network",
	DefaultAttestationLayer:       ATT_HOLESKY,
	AttestationLayer: map[string]*PresetAttestationConfig{
		ATT_HOLESKY: {
			Address: common.HexToAddress("0x2E8628F6000Ef85dea615af6Da4Fd6dF4fD149e6"),
		},
		ATT_AUTOMATA_TEST: {
			URL:     "https://rpc.ata.network",
			Address: common.HexToAddress("0xC9D1Fe39aC6259e66B3Be0e9DE5b33F8bbCa350F"),
		},
	},
}

var MainnetPreset = &PresetConfig{
	RegistryCoordinatorAddress:    common.HexToAddress("0x414696E4F7f06273973E89bfD3499e8666D63Bd4"),
	OperatorStateRetrieverAddress: common.HexToAddress("0x91246253d3Bff9Ae19065A90dC3AB6e09EefD2B6"),
	LineaProverURL:                "https://avs-prover-mainnet1.ata.network:18232",
	ScrollProverURL:               "https://avs-prover-mainnet1.ata.network:18232",
	DefaultAttestationLayer:       ATT_OPTIMISM,
	AttestationLayer: map[string]*PresetAttestationConfig{
		ATT_OPTIMISM: {
			URL:     "",
			Address: common.HexToAddress("0x99886d5C39c0DF3B0EAB67FcBb4CA230EF373510"),
		},
		ATT_AUTOMATA: {
			URL:     "https://rpc.ata.network",
			Address: common.HexToAddress("0xC9D1Fe39aC6259e66B3Be0e9DE5b33F8bbCa350F"),
		},
	},
}

func PresetConfigByRegistryCoordinatorAddress(registryCoordinatorAddress common.Address) *PresetConfig {
	for _, preset := range PresetConfigs {
		if preset.RegistryCoordinatorAddress == registryCoordinatorAddress {
			return preset
		}
	}
	return nil
}
