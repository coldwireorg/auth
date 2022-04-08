package auth

import (
	"auth/utils"
	"log"

	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	challenge := c.Query("logout_challenge")

	redirect, err := cwhydra.LogoutManager(*cwhydra.AdminApi).Accept(challenge)
	if err != nil {
		log.Println(err)
	}

	utils.DelCookie(c, "id_token")
	utils.DelCookie(c, "access_token")

	return c.Redirect(redirect)
}
