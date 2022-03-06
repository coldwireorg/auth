package controller

import (
	"auth/oauth2"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Callback(c *fiber.Ctx) error {
	code := c.Query("code")

	claims := oauth2.Callback(code)
	log.Println(claims)

	return c.Redirect("/user")
}
