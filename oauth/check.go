package oauth

import (
	"context"
)

// Check token validity
func CheckToken(idToken string, accessToken string) bool {
	idt, err := Verifier.Verify(context.Background(), idToken)
	if err != nil {
		return false
	}

	err = idt.VerifyAccessToken(accessToken)
	if err != nil {
		return false
	}

	return true
}
