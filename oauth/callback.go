package oauth

import (
	"log"
)

func Callback(code string) (string, string) {

	if code == "" {
		return "", ""
	}

	oauth2Token, err := Config.Exchange(Context, code)
	if err != nil {
		log.Println(err)
	}

	if oauth2Token == nil {
		return "", ""
	}

	// Extract the ID Token from OAuth2 token.
	idToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return "", ""
	}

	accessToken, ok := oauth2Token.Extra("access_token").(string)
	if !ok {
		return "", ""
	}

	return idToken, accessToken
}
