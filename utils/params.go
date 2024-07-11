package utils

import "github.com/ethereum/go-ethereum/common"

type PresetConfig struct {
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
	TEELivenessVerifierAddress    common.Address
	LineaProverURL                string
	ScrollProverURL               string
}

var PresetConfigs = []*PresetConfig{HoleskyTestnetPreset, MainnetPreset}

var HoleskyTestnetPreset = &PresetConfig{
	RegistryCoordinatorAddress:    common.HexToAddress("0x62c715575cE3Ad7C5a43aA325b881c70564f2215"),
	OperatorStateRetrieverAddress: common.HexToAddress("0xbfd43ac0a19c843e44491c3207ea13914818E214"),
	TEELivenessVerifierAddress:    common.HexToAddress("0x2E8628F6000Ef85dea615af6Da4Fd6dF4fD149e6"),
	LineaProverURL:                "https://avs-prover-staging.ata.network",
	ScrollProverURL:               "https://avs-prover-staging.ata.network",
}

var MainnetPreset = &PresetConfig{
	RegistryCoordinatorAddress:    common.HexToAddress("0x414696E4F7f06273973E89bfD3499e8666D63Bd4"),
	OperatorStateRetrieverAddress: common.HexToAddress("0x91246253d3Bff9Ae19065A90dC3AB6e09EefD2B6"),
	TEELivenessVerifierAddress:    common.HexToAddress("0x99886d5C39c0DF3B0EAB67FcBb4CA230EF373510"),
	LineaProverURL:                "https://avs-prover-mainnet1.ata.network:18232",
	ScrollProverURL:               "https://avs-prover-mainnet1.ata.network:18232",
}

func PresetConfigByRegistryCoordinatorAddress(registryCoordinatorAddress common.Address) *PresetConfig {
	for _, preset := range PresetConfigs {
		if preset.RegistryCoordinatorAddress == registryCoordinatorAddress {
			return preset
		}
	}
	return nil
}
