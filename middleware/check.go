package middleware

import (
	"auth/utils/errors"
	"auth/utils/tokens"

	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
)

// Verify if the user is already logged in
func OnAuth(c *fiber.Ctx) error {
	token := c.Cookies("token")
	challenge := c.Query("login_challenge")

	verifiedToken, err := tokens.Verify(token)
	if err != nil {
		return c.Next()
	}

	if challenge != "" {
		var claims tokens.Token
		verifiedToken.Claims(&claims)

		redirect, err := cwhydra.LoginManager(*cwhydra.AdminApi).Accept(challenge, cwhydra.AcceptLoginRequest{
			Subject: claims.Username,
			Context: map[string]interface{}{
				"username":    claims.Username,
				"role":        claims.Group,
				"private_key": claims.PrivateKey,
				"public_key":  claims.PublicKey,
			},
			Remember: true,
		})
		if err != nil {
			return errors.Handle(c, errors.ErrAuth, err)
		}

		return c.Redirect(redirect)
	} else {
		return c.Next()
	}
}

func CheckToken(c *fiber.Ctx) error {
	token := c.Cookies("token")

	_, err := tokens.Verify(token)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	return c.Next()
}
