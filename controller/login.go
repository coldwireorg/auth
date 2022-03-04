package controller

import (
	"auth/models"
	"auth/utils/errors"
	"context"

	"auth/hydra"

	h "github.com/ory/hydra-client-go"

	"github.com/alexedwards/argon2id"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		Username: username,
	}

	// Verify that the username is not empty
	if username == "" {
		return errors.HandleError(c, errors.ErrRequest, "sign-in")
	}

	// Get the user
	user, err := user.Get()
	if err != nil {
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

	// Decrypt private key with password
	// Now sending the encrypted one, decryption will happen on the client side :)
	/*	privateKey, err := chacha20.Decrypt(user.PrivateKey, cwcrypto.DeriveKey([]byte(password)))
		if err != nil {
			return errors.HandleError(c, errors.ErrInternal, "sign-in")
		} */

	acceptLoginRequest := h.NewAcceptLoginRequest(username)

	acceptLoginRequest.SetRemember(true)
	acceptLoginRequest.SetContext(fiber.Map{
		"username": username,
		"pvkey":    hexutil.Encode(user.PrivateKey),
	})

	resp, _, err := hydra.HydraAdminClient.AdminApi.AcceptLoginRequest(context.Background()).LoginChallenge(challenge).AcceptLoginRequest(*acceptLoginRequest).Execute()
	if err != nil {
		return err
	}

	return c.Redirect(resp.RedirectTo)
}
