package auth

import (
	"auth/utils"
	"auth/utils/errors"

	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	challenge := c.Query("logout_challenge")

	utils.DelCookie(c, "token")

	if challenge == "" {
		if c.Method() == "GET" {
			return c.Redirect("/")
		}

		return errors.Handle(c, errors.Success)
	} else {
		redirect, err := cwhydra.LogoutManager(*cwhydra.AdminApi).Accept(challenge)
		if err != nil {
			return errors.Handle(c, errors.ErrUnknown, err)
		}

		return c.Redirect(redirect)
	}
}
