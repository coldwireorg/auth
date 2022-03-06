package middleware

import (
	"auth/oauth2"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Check(c *fiber.Ctx) error {
	challenge := c.Query("login_challenge")

	if challenge != "" {
		return c.Next()
	}

	url, err := oauth2.AuthURL()
	if err != nil {
		log.Println(err)
	}

	return c.Redirect(url)
}
