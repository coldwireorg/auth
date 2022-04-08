package auth

import (
	"auth/models"
	"auth/utils/errors"

	"codeberg.org/coldwire/cwhydra"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	challenge := c.Query("login_challenge")

	if challenge == "" {
		return fiber.ErrBadRequest
	}

	user := models.User{
		Name: username,
	}

	// Verify that the username is not empty
	if username == "" {
		return errors.Handle(c, errors.ErrBody)
	}

	// Get the user
	user, err := user.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	if user.Name == "" {
		return errors.Handle(c, errors.ErrBody)
	}

	// Verify password
	isValidPassword, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)

	}

	if !isValidPassword {
		return errors.Handle(c, errors.ErrAuthPassword)
	}

	redirect, err := cwhydra.LoginManager(*cwhydra.AdminApi).Accept(challenge, cwhydra.AcceptLoginRequest{
		Subject: username,
		Context: map[string]interface{}{
			"username":    username,
			"role":        user.Group,
			"private_key": user.PrivateKey,
			"public_key":  user.PublicKey,
		},
		Remember: true,
	})
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	return c.Redirect(redirect)
}
