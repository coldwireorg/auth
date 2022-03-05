package controller

import (
	"auth/hydra"
	"auth/models"

	"auth/utils/errors"
	"context"
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	h "github.com/ory/hydra-client-go"
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
		Username: username,
	}

	// Verify that the username is not empty or too short
	if username == "" {
		return errors.HandleError(c, errors.ErrRequest, "sign-up")
	}

	// Verify if the user already exist
	exist := user.Exist()
	log.Println(exist)
	if exist {
		return errors.HandleError(c, errors.ErrAuthExist, "sign-up")
	}

	// Hash password with argon2id
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return errors.HandleError(c, errors.ErrInternal, "sign-up")
	}

	// Set user data
	user.Password = hash
	user.PrivateKey = pvkey
	user.PublicKey = pbkey

	err = user.Create()
	if err != nil {
		return errors.HandleError(c, errors.ErrDatabaseCreate, "sign-up")
	}

	acceptLoginRequest := h.NewAcceptLoginRequest(username)

	acceptLoginRequest.SetRemember(true)
	acceptLoginRequest.SetContext(fiber.Map{
		"username": username,
		"pvkey":    user.PrivateKey,
		"pbkey":    user.PublicKey,
	})

	resp, _, err := hydra.HydraAdminClient.AdminApi.AcceptLoginRequest(context.Background()).LoginChallenge(challenge).AcceptLoginRequest(*acceptLoginRequest).Execute()
	if err != nil {
		return err
	}

	return c.Redirect(resp.RedirectTo)
}
