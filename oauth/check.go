package oauth

// Check token validity
func CheckToken(idToken string, accessToken string) bool {
	idt, err := Verifier.Verify(Context, idToken)
	if err != nil {
		return false
	}

	err = idt.VerifyAccessToken(accessToken)
	if err != nil {
		return false
	}

	return true
}
