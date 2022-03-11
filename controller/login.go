package controller

import (
	"auth/models"
	"auth/utils/errors"
	"log"

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
		return errors.HandleError(c, errors.ErrRequest, "sign-in")
	}

	// Get the user
	user, err := user.Find()
	if err != nil {
		return errors.HandleError(c, errors.ErrDatabaseNotFound, "sign-in")
	}

	if user.Name == "" {
		return errors.HandleError(c, errors.ErrDatabaseNotFound, "sign-in")
	}

	// Verify password
	isValidPassword, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return errors.HandleError(c, errors.ErrInternal, "sign-in")
	}

	if !isValidPassword {
		return errors.HandleError(c, errors.ErrAuthPassword, "sign-in")
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
		return errors.HandleError(c, errors.ErrInternal, "sign-in")
	}

	/*
		ctx, err := json.Marshal(map[string]interface{}{
			"username":    username,
			"role":        user.Role,
			"private_key": user.PrivateKey,
			"public_key":  user.PublicKey,
		})
		if err != nil {
			log.Println(err)
		}

		utils.SetCookie(c, "ctx", string(ctx), time.Now().Add(time.Second*30))
	*/

	log.Println(redirect)

	return c.Redirect(redirect)
}
