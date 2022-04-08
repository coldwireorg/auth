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
		return errors.Handle(c, errors.ErrBody)
	}

	// Verify if the user already exist
	exist := user.Exist()
	if exist {
		return errors.Handle(c, errors.ErrAuthExist)
	}

	// Hash password with argon2id
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	// If this is the first user to register, makes them admin
	isFirstUser := user.IsFirstOne()

	if isFirstUser {
		user.Group = "admin"
	} else {
		user.Group = "user"
	}

	// Set user data
	user.Password = hash
	user.PrivateKey = pvkey
	user.PublicKey = pbkey

	err = user.Create()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
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
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	return c.Redirect(redirect)
}
