package utils

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

func ExpandPubkey(pubkey *bls.G1Point) ([32]byte, [32]byte) {
	pubkeyBytes := pubkey.Serialize()
	if len(pubkeyBytes) != 64 {
		panic("invalid pubkey")
	}
	var x, y [32]byte
	copy(x[:], pubkeyBytes[:32])
	copy(y[:], pubkeyBytes[32:64])
	return x, y
}

func SplitPubkey(pubkey []byte) ([32]byte, [32]byte) {
	if len(pubkey) != 64 {
		panic("invalid pubkey")
	}
	var x, y [32]byte
	copy(x[:], pubkey[:32])
	copy(y[:], pubkey[32:64])
	return x, y
}
