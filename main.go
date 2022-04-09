package main

import (
	"auth/database"
	"auth/routes"
	"auth/utils"
	"auth/utils/config"
	"auth/utils/env"
	"auth/utils/tokens"
	"crypto/tls"
	"embed"
	"flag"
	"net/http"
	"os"

	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed views/public/*
var views embed.FS

func init() {
	// Configure logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fileConf := flag.String("config", "", "Path to the config file")
	flag.Parse()

	// Init configuration
	config.Init(env.Get("CONFIG_FILE", *fileConf))

	// Connect to database
	database.Connect()

	// Connect to hydra admin api
	cwhydra.Init(cwhydra.Config{
		Url: config.Conf.Hydra.Admin,
		Http: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 0,
		},
	})

	// generate key for signing jwt tokens
	tokens.Init()
}

func main() {
	// Create fiber instance
	app := fiber.New()

	// migrate database
	utils.MigrateTables()

	// Include cors
	app.Use(cors.New())

	// Setup routes
	routes.Api(app)

	if config.Conf.Hydra.Proxy == "true" {
		routes.Hydra(app)
	}

	// Load view as static website
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(views),
		PathPrefix: "views/public",
		Browse:     true,
	}))

	utils.InitClients()

	log.Info().Err(app.Listen(config.Conf.Server.Address + ":" + config.Conf.Server.Port))
}
