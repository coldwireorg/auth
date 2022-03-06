package main

import (
	"auth/database"
	"auth/hydra"
	"auth/oauth"
	"auth/router"
	"log"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
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

	go oauth.InitOauth2(oauth2.Config{
		ClientID:    "auth",
		RedirectURL: os.Getenv("AUTH_SERVER_URL") + "/api/callback",
		Scopes:      []string{oidc.ScopeOpenID, "profile"},
	})

	go hydra.NewHydraClient(os.Getenv("HYDRA_ADMIN_URL"))
	go database.Connect()

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
