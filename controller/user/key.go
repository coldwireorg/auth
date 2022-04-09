package user

import (
	"auth/models"
	"auth/utils/errors"

	"github.com/gofiber/fiber/v2"
)

func Key(c *fiber.Ctx) error {
	username := c.Params("username")

	user := models.User{
		Name: username,
	}

	key, err := user.Pubkey()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)

	}

	return errors.Handle(c, errors.Success, fiber.Map{
		"key": key,
	})
}
