package cwcrypto

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func GenerateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	pvkey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	pbkey := &pvkey.PublicKey

	return pvkey, pbkey, nil
}

// Encode public and private keys into bytes
func Encode(pvkey *ecdsa.PrivateKey, pbkey *ecdsa.PublicKey) ([]byte, []byte, error) {
	vkey := crypto.FromECDSA(pvkey)
	bkey := crypto.FromECDSAPub(pbkey)

	return vkey, bkey, nil
}

// Decode public and private keyz into structs
func Decode(pvkey []byte, pbkey []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	vkey, err := crypto.ToECDSA(pvkey)
	bkey, err := crypto.UnmarshalPubkey(pbkey)
	if err != nil {
		return nil, nil, err
	}

	return vkey, bkey, nil
}
