package router

import (
	"auth/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupApi(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Get("/consent", controller.Consent)
	api.Get("/callback", controller.Callback)

	api.Get("/pubkey/:username?", controller.Pubkey)
}
