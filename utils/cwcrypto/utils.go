package cwcrypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"
)

// Derivate a key or a password with argon2
func DeriveKey(key []byte) []byte {
	// Generating derivated key from a byte array
	// this function use the recomanded settings for
	// generating a 32 bytes key.
	key = argon2.Key(key, nil, 3, 32*1024, 4, 32)
	return key // return key
}

func Sign(pvkey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	sig, err := secp256k1.Sign(data, pvkey.X.Bytes())
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func Hash(data []byte) []byte {
	algo := sha3.New256()
	algo.Write(data)
	return algo.Sum(nil)
}
