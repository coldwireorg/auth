package auth

import (
	"auth/utils/errors"

	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
)

func Consent(c *fiber.Ctx) error {
	challenge := c.Query("consent_challenge")

	login, err := cwhydra.ConsentManager(*cwhydra.AdminApi).Get(challenge)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	redirect, err := cwhydra.ConsentManager(*cwhydra.AdminApi).Accept(challenge, cwhydra.AcceptConsentRequest{
		GrantScope: []string{
			"openid",
			"offline",
		},
		Session: cwhydra.ConsentRequestSession{
			IdToken: login.Context,
		},
		Remember: true,
	})

	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	return c.Redirect(redirect)
}
