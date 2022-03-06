package controller

import (
	"auth/hydra"
	"auth/utils"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Consent(c *fiber.Ctx) error {
	challenge := c.Query("consent_challenge")
	ctx := c.Cookies("ctx")

	var payload map[string]interface{}
	err := json.Unmarshal([]byte(ctx), &payload)
	if err != nil {
		log.Println(err)
	}

	log.Println(payload)

	redirect, err := hydra.Consent(challenge, payload)
	if err != nil {
		log.Println(err)
		c.Redirect(redirect)
	}

	utils.DelCookie(c, "ctx")
	return c.Redirect(redirect)
}
