package errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type apiError struct {
	status  int
	code    string
	message string
}

func HandleError(c *fiber.Ctx, err apiError, path string) error {
	log.Error().Msg(err.message)
	return c.Redirect("/" + path + "?err=" + err.message)
}
