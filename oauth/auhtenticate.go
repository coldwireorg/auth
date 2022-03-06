package oauth

import (
	"auth/utils"

	"github.com/coreos/go-oidc/v3/oidc"
)

func AuthURL() (string, error) {
	state, err := utils.GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	nonce, err := utils.GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	return Config.AuthCodeURL(state, oidc.Nonce(nonce)), nil
}
