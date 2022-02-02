package errors

import "github.com/gofiber/fiber/v2"

type apiError struct {
	status  int
	code    string
	message string
}

func HandleError(c *fiber.Ctx, err apiError, path string) error {
	return c.Redirect("/" + path + "?err=" + err.message)
}
