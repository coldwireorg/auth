package oauth

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
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
		ctx := oidc.ClientContext(context.Background(), &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 0,
		})

		Provider, err = oidc.NewProvider(ctx, os.Getenv("HYDRA_PUBLIC_URL"))
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
