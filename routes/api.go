package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	api.Use(logger.New())

	auth := api.Group("/auth")
	auth.All("/register") // register
	auth.All("/login")    // login
	auth.All("/consent")  // consent oauth request
	auth.All("/logout")   // logout

	user := api.Group("/user")
	user.Get("/key/:username") // get user's public key
	user.Put("/")              // update user's informations
	user.Delete("/")           // delte user

	services := api.Group("/services")
	services.Get("/")       // list all services
	services.Post("/:id")   // add a new service
	services.Put("/:id")    // update a service
	services.Delete("/:id") // delete a service
}
