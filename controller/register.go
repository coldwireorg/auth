package controller

import (
	"auth/models"
	"auth/utils"
	"auth/utils/cwcrypto"
	"auth/utils/cwcrypto/chacha20"
	"auth/utils/errors"
	"auth/utils/tokens"
	"log"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.User{
		Username: username,
	}

	// Verify that the username is not empty or too short
	if username == "" {
		return errors.HandleError(c, errors.ErrRequest)
	}

	// Verify if the user already exist
	exist := user.Exist()
	log.Println(exist)
	if exist {
		return errors.HandleError(c, errors.ErrAuthExist)
	}

	// Hash password with argon2id
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return errors.HandleError(c, errors.ErrInternal)
	}

	// Generate user's keypair
	pvkey, pbkey, err := cwcrypto.GenerateKeys()
	if err != nil {
		return c.SendString(err.Error())
	}

	// Encode keys to store them
	encodedPvkey, encodedPbkey, err := cwcrypto.Encode(pvkey, pbkey)
	if err != nil {
		return c.SendString(err.Error())
	}

	// Encrypt private key with user's password wich is derived using Argon2i
	pvKeyEncryptionKey := cwcrypto.DeriveKey([]byte(password))
	encryptedPvkey, err := chacha20.Encrypt(encodedPvkey, pvKeyEncryptionKey)
	if err != nil {
		return errors.HandleError(c, errors.ErrInternal)
	}

	// Set user data
	user.Password = hash
	user.PublicKey = encodedPbkey
	user.PrivateKey = encryptedPvkey

	err = user.Create()
	if err != nil {
		return errors.HandleError(c, errors.ErrDatabaseCreate)
	}

	exp := time.Hour * 2                                        // define token expiration
	jwt := tokens.Generate(username, string(encodedPvkey), exp) // Generate JWT token

	// Set cookies
	c.Cookie(utils.GenCookie("token", jwt, exp, os.Getenv("SERVER_DOMAIN")))
	// If the user use TOR, set the cookies on the tor address
	if c.Hostname() == os.Getenv("TOR_ADDRESS") {
		c.Cookie(utils.GenCookie("token", jwt, exp, os.Getenv("TOR_ADDRESS")))
	}

	// Return login informations
	return c.Redirect("/user/" + username)
}
