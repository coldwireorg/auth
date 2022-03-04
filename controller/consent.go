package controller

import (
	"auth/hydra"
	"context"

	"github.com/gofiber/fiber/v2"
	h "github.com/ory/hydra-client-go"
)

func Consent(c *fiber.Ctx) error {
	challenge := c.Query("consent_challenge")

	acceptConsentRequest := h.NewAcceptConsentRequest()

	acceptConsentRequest.SetRemember(true)
	acceptConsentRequest.SetGrantScope([]string{
		"openid",
		"offline",
	})

	resp, _, err := hydra.HydraAdminClient.AdminApi.AcceptConsentRequest(context.Background()).ConsentChallenge(challenge).AcceptConsentRequest(*acceptConsentRequest).Execute()
	if err != nil {
		return err
	}

	return c.Redirect(resp.RedirectTo)
}
