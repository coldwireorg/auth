package routes

import (
	"auth/controller/auth"
	"auth/controller/user"
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
	authentication.All("/consent", auth.Consent)                      // consent oauth request
	authentication.All("/logout", auth.Logout)                        // logout

	usr := api.Group("/user")
	usr.Get("/key/:username", user.Key)                        // get user's public key
	usr.Put("/password", middleware.CheckToken, user.Password) // update user's informations
	usr.Delete("/", middleware.CheckToken, user.Delete)        // delete user

	/*
		services := api.Group("/services")
		services.Get("/")       // list all services
		services.Post("/:id")   // add a new service
		services.Put("/:id")    // update a service
		services.Delete("/:id") // delete a service
	*/
}
