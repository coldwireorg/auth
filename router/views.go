package router

import (
	"auth/middleware"
	"auth/oauth"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupViews(app *fiber.App) {
	app.Get("/", middleware.CheckAuthenticated, func(c *fiber.Ctx) error {
		return c.Redirect("/sign-in")
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		return c.Redirect(oauth.LogoutURL())
	})

	app.Get("/sign-in", middleware.CheckAuthenticated, func(c *fiber.Ctx) error {
		err := c.Query("err")
		challenge := c.Query("login_challenge")

		log.Println(err)

		return c.Render("index", fiber.Map{
			"Register":  false,
			"Challenge": challenge,
			"Error":     err,
		}, "layouts/auth")
	})

	app.Get("/sign-up", middleware.CheckAuthenticated, func(c *fiber.Ctx) error {
		err := c.Query("err")
		challenge := c.Query("login_challenge")

		return c.Render("index", fiber.Map{
			"Register":  true,
			"Challenge": challenge,
			"Error":     err,
		}, "layouts/auth")
	})

	app.Get("/user", middleware.CheckUser, func(c *fiber.Ctx) error {
		return c.Render("board", fiber.Map{}, "layouts/user")
	})
}
