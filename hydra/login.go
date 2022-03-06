package hydra

import (
	"context"

	h "github.com/ory/hydra-client-go"
)

func Login(username string, challenge string, ctx map[string]interface{}) (string, error) {
	acceptLoginRequest := h.NewAcceptLoginRequest(username)

	acceptLoginRequest.SetRemember(true)
	acceptLoginRequest.SetContext(ctx)

	resp, _, err := Client.AdminApi.AcceptLoginRequest(context.Background()).LoginChallenge(challenge).AcceptLoginRequest(*acceptLoginRequest).Execute()
	if err != nil {
		return "/sign-in", err
	}

	return resp.RedirectTo, nil
}
