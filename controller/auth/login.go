package auth

import (
	"auth/models"
	"auth/utils"
	"auth/utils/errors"
	"auth/utils/tokens"
	"time"

	"codeberg.org/coldwire/cwhydra"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	challenge := c.Query("login_challenge")

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

	usrData := tokens.Token{
		Username:   user.Name,
		Group:      user.Group,
		PrivateKey: user.PrivateKey,
		PublicKey:  user.PublicKey,
	}

	cookie := tokens.Generate(usrData, 24*time.Hour)
	utils.SetCookie(c, "token", cookie, time.Now())

	if challenge == "" {
		return errors.Handle(c, errors.Success, usrData)
	} else {
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
}
