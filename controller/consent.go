package controller

import (
	"auth/hydra"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Consent(c *fiber.Ctx) error {
	challenge := c.Query("consent_challenge")

	redirect, err := hydra.Consent(challenge)
	if err != nil {
		log.Println(err)
		c.Redirect(redirect)
	}

	return c.Redirect(redirect)
}
