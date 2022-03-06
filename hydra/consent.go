package hydra

import (
	"context"

	h "github.com/ory/hydra-client-go"
)

func Consent(challenge string, payload map[string]interface{}) (string, error) {
	acceptConsentRequest := h.NewAcceptConsentRequest()

	acceptConsentRequest.SetRemember(true)
	acceptConsentRequest.SetGrantScope([]string{
		"openid",
		"offline",
	})

	consentRequestSession := h.NewConsentRequestSession()
	consentRequestSession.SetIdToken(payload)
	acceptConsentRequest.SetSession(*consentRequestSession)

	resp, _, err := Client.AdminApi.AcceptConsentRequest(context.Background()).ConsentChallenge(challenge).AcceptConsentRequest(*acceptConsentRequest).Execute()
	if err != nil {
		return "/sign-in", err
	}

	return resp.RedirectTo, nil
}
