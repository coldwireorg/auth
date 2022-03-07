package oauth

import (
	"log"
)

type Claims struct {
	Username   string `json:"username"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// Check token validity
func GetClaims(token string) Claims {
	// Parse and verify ID Token payload.
	idToken, err := Verifier.Verify(Context, token)
	if err != nil {
		log.Println(err)
	}

	// Extract custom claims
	var claims Claims
	err = idToken.Claims(&claims)
	if err != nil {
		log.Println(err)
	}

	return claims
}
