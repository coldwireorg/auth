package controller

import (
	"auth/hydra"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	h "github.com/ory/hydra-client-go"
)

func Consent(c *fiber.Ctx) error {
	challenge := c.Query("consent_challenge")

	acceptConsentRequest := h.NewAcceptConsentRequest()

	ctx := h.NewOpenIDConnectContext()
	log.Println(ctx.Display)

	acceptConsentRequest.SetRemember(true)
	acceptConsentRequest.SetGrantScope([]string{
		"openid",
		"offline",
	})

	consentRequestSession := h.NewConsentRequestSession()

	acceptConsentRequest.SetSession(*consentRequestSession)

	resp, _, err := hydra.HydraAdminClient.AdminApi.AcceptConsentRequest(context.Background()).ConsentChallenge(challenge).AcceptConsentRequest(*acceptConsentRequest).Execute()
	if err != nil {
		return err
	}

	return c.Redirect(resp.RedirectTo)
}
