package oauth

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	Provider *oidc.Provider
	Verifier *oidc.IDTokenVerifier
	Config   *oauth2.Config
)

func InitOauth2(conf oauth2.Config) error {
	var err error

	for {
		Provider, err = oidc.NewProvider(context.Background(), os.Getenv("HYDRA_PUBLIC_URL"))
		if err != nil {
			log.Println(err)
			time.Sleep(15 * time.Second)
		} else {
			conf.Endpoint = Provider.Endpoint()
			Config = &conf
			Verifier = Provider.Verifier(&oidc.Config{ClientID: Config.ClientID})
			log.Println("Successfully connected to the oauth2 server!")
			break
		}
	}

	return nil
}
