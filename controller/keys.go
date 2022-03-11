package controller

import (
	"auth/models"

	"github.com/gofiber/fiber/v2"
)

func Pubkey(c *fiber.Ctx) error {
	username := c.Params("username")

	user := models.User{
		Name: username,
	}

	key, err := user.Pubkey()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"key": key,
	})
}
