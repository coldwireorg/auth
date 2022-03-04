package hydra

import (
	"crypto/tls"
	"net/http"

	hydra "github.com/ory/hydra-client-go"
)

var (
	HydraAdminURL    = "127.0.0.1:4445"
	HydraAdminClient = NewHydraClient(HydraAdminURL)
)

func NewHydraClient(endpoint string) *hydra.APIClient {
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

	return apiClient
}
