package oauth2

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	Provider *oidc.Provider
	Verifier *oidc.IDTokenVerifier
	Config   *oauth2.Config
)

func InitOauth2() {
	var err error

	Provider, err = oidc.NewProvider(context.Background(), os.Getenv("HYDRA_PUBLIC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	Config = &oauth2.Config{
		ClientID:    "auth",
		Endpoint:    Provider.Endpoint(),
		RedirectURL: "http://" + os.Getenv("DOMAIN") + "/api/callback",
		Scopes:      []string{oidc.ScopeOpenID, "profile"},
	}

	Verifier = Provider.Verifier(&oidc.Config{ClientID: Config.ClientID})
}
