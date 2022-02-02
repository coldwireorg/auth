package controller

import (
	"auth/models"
	"auth/utils"
	"auth/utils/cwcrypto"
	"auth/utils/cwcrypto/chacha20"
	"auth/utils/errors"
	"auth/utils/tokens"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

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
	privateKey, err := chacha20.Decrypt(user.PrivateKey, cwcrypto.DeriveKey([]byte(password)))
	if err != nil {
		return errors.HandleError(c, errors.ErrInternal, "sign-in")
	}

	exp := time.Hour * 2
	jwt := tokens.Generate(username, hexutil.Encode(privateKey), exp) // Generate JWT token

	// set cookies
	c.Cookie(utils.GenCookie("token", jwt, exp, os.Getenv("SERVER_DOMAIN")))
	// if user use tor browser
	if c.Hostname() == os.Getenv("TOR_ADDRESS") {
		c.Cookie(utils.GenCookie("token", jwt, exp, os.Getenv("TOR_ADDRESS")))
	}

	// Return login informations
	return c.Redirect("/user/" + username)
}
