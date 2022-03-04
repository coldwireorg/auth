package middleware

import (
	"github.com/gofiber/fiber/v2"
	h "github.com/ory/hydra-client-go"
)

func CheckUser(c *fiber.Ctx) error {
	res := h.NewUserinfoResponse()
	return c.JSON(fiber.Map{
		"res": res,
	})
}
