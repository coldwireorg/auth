package router

import (
	"auth/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupViews(app *fiber.App) {
	app.Get("/", middleware.Check, func(c *fiber.Ctx) error {
		return c.Redirect("/sign-in")
	})

	app.Get("/sign-in", middleware.Check, func(c *fiber.Ctx) error {
		err := c.Query("err")
		challenge := c.Query("login_challenge")

		log.Println(err)

		return c.Render("index", fiber.Map{
			"Register":  false,
			"Challenge": challenge,
			"Error":     err,
		}, "layouts/auth")
	})

	app.Get("/sign-up", middleware.Check, func(c *fiber.Ctx) error {
		err := c.Query("err")
		challenge := c.Query("login_challenge")

		return c.Render("index", fiber.Map{
			"Register":  true,
			"Challenge": challenge,
			"Error":     err,
		}, "layouts/auth")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.Render("board", fiber.Map{}, "layouts/user")
	})
}
