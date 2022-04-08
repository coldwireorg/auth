package routes

import (
	"auth/controller/auth"
	"auth/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	api.Use(logger.New())

	authentication := api.Group("/auth")
	authentication.All("/register", middleware.OnAuth, auth.Register) // register
	authentication.All("/login", middleware.OnAuth, auth.Login)       // login
	//authentication.All("/consent")  // consent oauth request
	authentication.All("/logout", auth.Logout) // logout

	/*
		user := api.Group("/user")
		user.Get("/key/:username") // get user's public key
		user.Put("/")              // update user's informations
		user.Delete("/")           // delte user

		services := api.Group("/services")
		services.Get("/")       // list all services
		services.Post("/:id")   // add a new service
		services.Put("/:id")    // update a service
		services.Delete("/:id") // delete a service
	*/
}
