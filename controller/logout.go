package controller

import (
	"auth/hydra"
	"auth/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	challenge := c.Query("logout_challenge")

	redirect, err := hydra.Logout(challenge)
	if err != nil {
		log.Println(err)
	}

	utils.DelCookie(c, "id_token")
	utils.DelCookie(c, "access_token")

	return c.Redirect(redirect)
}
