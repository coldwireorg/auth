package controller

import (
	"auth/models"
	"auth/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Pubkey(c *fiber.Ctx) error {
	username := c.Params("username")

	if username != "" {
		user := models.User{
			Username: username,
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
	} else {
		return c.JSON(fiber.Map{
			"key": tokens.JWTPrivateKey.Public(),
		})
	}
}
