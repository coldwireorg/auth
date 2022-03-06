package oauth2

import (
	"context"
	"log"
)

type Claims struct {
	Username   string `json:"username"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func Callback(code string) Claims {

	if code == "" {
		return Claims{}
	}

	oauth2Token, err := Config.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
	}

	if oauth2Token == nil {
		return Claims{}
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return Claims{}
	}

	// Parse and verify ID Token payload.
	idToken, err := Verifier.Verify(context.Background(), rawIDToken)
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
