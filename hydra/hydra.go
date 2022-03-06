package hydra

import (
	"crypto/tls"
	"net/http"

	hydra "github.com/ory/hydra-client-go"
)

var Client *hydra.APIClient

func NewHydraClient(endpoint string) {
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

	Client = apiClient
}
