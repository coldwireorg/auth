package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupViews(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/sign-in")
	})

	app.Get("/sign-in", func(c *fiber.Ctx) error {
		err := c.Query("err")

		log.Println(err)

		return c.Render("index", fiber.Map{
			"Register": false,
			"Error":    err,
		}, "layouts/auth")
	})

	app.Get("/sign-up", func(c *fiber.Ctx) error {
		err := c.Query("err")

		return c.Render("index", fiber.Map{
			"Register": true,
			"Error":    err,
		}, "layouts/auth")
	})
}
