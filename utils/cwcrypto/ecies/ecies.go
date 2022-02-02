package ecies

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func Encrypt(pbkey *ecdsa.PublicKey, data []byte) ([]byte, error) {
	public := ecies.ImportECDSAPublic(pbkey)

	encrypted, err := ecies.Encrypt(rand.Reader, public, data, nil, nil)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func Decrypt(pvkey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	private := ecies.ImportECDSA(pvkey)

	decrypted, err := private.Decrypt(data, nil, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
