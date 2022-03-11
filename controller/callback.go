package controller

import (
	"auth/utils"
	"time"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
)

func Callback(c *fiber.Ctx) error {
	code := c.Query("code")

	idToken, accessToken := cwauth.Callback(code)

	utils.SetCookie(c, "id_token", idToken, time.Now().Add(time.Hour*6))
	utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*6))

	return c.Redirect("/user")
}
