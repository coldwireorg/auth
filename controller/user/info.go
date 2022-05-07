package user

import (
	"auth/models"
	"auth/utils/errors"
	"auth/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {
	token := c.Cookies("token")

	t, err := tokens.Parse(token)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	usr := models.User{
		Name: t.Username,
	}

	usr, err = usr.Find()

	return errors.Handle(c, errors.Success, usr)
}
