package controller

import (
	"auth/models"

	"auth/utils/errors"

	"codeberg.org/coldwire/cwhydra"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	pvkey := c.FormValue("pvkey")
	pbkey := c.FormValue("pbkey")

	challenge := c.Query("login_challenge")

	if challenge == "" {
		return fiber.ErrBadRequest
	}

	user := models.User{
		Name: username,
	}

	// Verify that the username is not empty or too short
	if username == "" {
		return errors.HandleError(c, errors.ErrRequest, "sign-up")
	}

	// Verify if the user already exist
	exist := user.Exist()
	if exist {
		return errors.HandleError(c, errors.ErrAuthExist, "sign-up")
	}

	// Hash password with argon2id
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return errors.HandleError(c, errors.ErrInternal, "sign-up")
	}

	// Set user data
	user.Group = "user"
	user.Password = hash
	user.PrivateKey = pvkey
	user.PublicKey = pbkey

	err = user.Create()
	if err != nil {
		return errors.HandleError(c, errors.ErrDatabaseCreate, "sign-up")
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
		return errors.HandleError(c, errors.ErrInternal, "sign-up")
	}

	return c.Redirect(redirect)
}
