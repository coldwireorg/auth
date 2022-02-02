package main

import (
	"auth/database"
	"auth/router"
	"auth/utils/tokens"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load env values: %v\n", err)
	} else {
		log.Println("Loaded env values successfully")
	}
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.Connect()

	// Generate JWT signin keys
	tokens.GenerateKeys()

	app.Use(cors.New()) // Add cors
	app.Static("/static", "./static")

	router.SetupViews(app)
	router.SetupApi(app)

	listener := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	err := app.Listen(listener)
	if err != nil {
		log.Println("Unable to start server on [" + listener + "]")
		panic(err)
	} else {
		log.Println("Listening on [" + listener + "]")
	}
}
