package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupViews(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/sign-in")
	})

	app.Get("/sign-in", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Register": false,
			"Error":    false,
		}, "layouts/auth")
	})

	app.Get("/sign-up", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Register": true,
		}, "layouts/auth")
	})
}
