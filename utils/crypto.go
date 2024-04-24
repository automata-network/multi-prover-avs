package utils

import (
	"crypto/ecdsa"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ZeroAddress common.Address
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

func EcdsaAddress(key *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(*key.Public().(*ecdsa.PublicKey))
}
