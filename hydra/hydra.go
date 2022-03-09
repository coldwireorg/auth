package hydra

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	hydra "github.com/ory/hydra-client-go"
)

var (
	Client *hydra.APIClient
)

func Connect(endpoint string) bool {
	configuration := hydra.NewConfiguration()

	configuration.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
	}

	apiClient := hydra.NewAPIClient(configuration)
	apiClient.GetConfig().Host = endpoint
	apiClient.GetConfig().Scheme = "http"

	for {
		resp, _, err := apiClient.MetadataApi.IsReady(context.Background()).ApiService.IsAlive(context.Background()).Execute()

		if resp.GetStatus() == "ok" && err == nil {
			log.Println("Connected to hydra!")
			Client = &hydra.APIClient{}
			return true
		} else {
			log.Println(err.Error())
		}
	}
}
