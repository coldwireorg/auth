package hydra

import (
	"context"

	h "github.com/ory/hydra-client-go"
)

func Consent(challenge string) (string, error) {
	acceptConsentRequest := h.NewAcceptConsentRequest()

	acceptConsentRequest.SetRemember(true)
	acceptConsentRequest.SetGrantScope([]string{
		"openid",
		"offline",
	})

	consentRequestSession := h.NewConsentRequestSession()
	consentRequestSession.SetIdToken(map[string]interface{}{
		"username":   "monoko",
		"email":      "monoko@coldwire.org",
		"privateKey": "privateKey",
		"publicKey":  "publicKey",
	})
	acceptConsentRequest.SetSession(*consentRequestSession)

	resp, _, err := Client.AdminApi.AcceptConsentRequest(context.Background()).ConsentChallenge(challenge).AcceptConsentRequest(*acceptConsentRequest).Execute()
	if err != nil {
		return "/sign-in", err
	}

	return resp.RedirectTo, nil
}
