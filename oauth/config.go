package oauth

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

func InitOauth2(conf oauth2.Config) {
	var err error

	Provider, err = oidc.NewProvider(context.Background(), os.Getenv("HYDRA_PUBLIC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	conf.Endpoint = Provider.Endpoint()

	Config = &conf

	Verifier = Provider.Verifier(&oidc.Config{ClientID: Config.ClientID})
}
