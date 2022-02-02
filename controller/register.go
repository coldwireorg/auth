package controller

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	return c.SendString("")
}
