package hydra

import (
	"crypto/tls"
	"net/http"
	"os"

	hydra "github.com/ory/hydra-client-go"
)

var (
	Client = NewHydraClient(os.Getenv("HYDRA_ADMIN_URL"))
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
