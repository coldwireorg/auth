package routes

import (
	"auth/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	app.Use(logger.New())

	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Get("/logout", controller.Logout)
	api.Get("/consent", controller.Consent)
	api.Get("/callback", controller.Callback)

	api.Get("/pubkey/:username", controller.Pubkey)
}
