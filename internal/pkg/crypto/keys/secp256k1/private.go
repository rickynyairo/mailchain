package secp256k1

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// PrivateKey based on the secp256k1 curve
type PrivateKey struct {
	ecdsa ecdsa.PrivateKey
}

// Bytes returns the byte representation of the private key
func (pk PrivateKey) Bytes() []byte {
	return crypto.FromECDSA(&pk.ecdsa)
}

// PrivateKeyFromECDSA get a private key from an ecdsa.PrivateKey
func PrivateKeyFromECDSA(pk ecdsa.PrivateKey) PrivateKey {
	return PrivateKey{ecdsa: pk}
}

// PrivateKeyFromBytes get a private key from []byte
func PrivateKeyFromBytes(pk []byte) (*PrivateKey, error) {
	rpk, err := crypto.ToECDSA(pk)
	if err != nil {
		return nil, errors.Errorf("could not convert private key")
	}
	return &PrivateKey{ecdsa: *rpk}, nil
}

// PrivateKeyFromHex get a private key from hex string
func PrivateKeyFromHex(hexkey string) (*PrivateKey, error) {
	b, err := hex.DecodeString(hexkey)
	if err != nil {
		return nil, errors.New("invalid hex string")
	}
	return PrivateKeyFromBytes(b)
}
