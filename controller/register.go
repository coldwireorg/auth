package controller

import (
	"auth/hydra"
	"auth/models"
	"auth/utils"
	"encoding/json"
	"time"

	"auth/utils/errors"
	"log"

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

	redirect, err := hydra.Login(username, challenge)
	if err != nil {
		log.Println(err)
	}

	ctx, err := json.Marshal(map[string]interface{}{
		"username": username,
		"pvkey":    user.PrivateKey,
		"pbkey":    user.PublicKey,
	})
	if err != nil {
		log.Println(err)
	}

	utils.SetCookie(c, "ctx", string(ctx), time.Now().Add(time.Second*30))

	return c.Redirect(redirect)
}
