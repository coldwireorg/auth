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
	challenge := c.Query("login_challenge")

	request := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		return errors.Handle(c, errors.ErrBody, err)
	}

	user := models.User{
		Name: request.Username,
	}

	// Verify that the username is not empty
	if request.Username == "" {
		return errors.Handle(c, errors.ErrBody)
	}

	// Get the user
	user, err = user.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	if user.Name == "" {
		return errors.Handle(c, errors.ErrDatabaseNotFound)
	}

	// Verify password
	isValidPassword, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
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
			Subject: request.Username,
			Context: map[string]interface{}{
				"username":    request.Username,
				"role":        user.Group,
				"private_key": user.PrivateKey,
				"public_key":  user.PublicKey,
			},
			Remember: true,
		})
		if err != nil {
			return errors.Handle(c, errors.ErrAuth, err)
		}

		return errors.Handle(c, errors.Success, fiber.Map{
			"redirect_url": redirect,
		})
	}
}
