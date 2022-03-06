package controller

import (
	"auth/oauth"
	"auth/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Callback(c *fiber.Ctx) error {
	code := c.Query("code")

	idToken, accessToken := oauth.Callback(code)

	utils.SetCookie(c, "id_token", idToken, time.Now().Add(time.Hour*6))
	utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*6))

	claims := oauth.GetClaims(idToken)
	log.Println(claims)

	return c.Redirect("/user")
}
