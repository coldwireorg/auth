package main

import (
	"auth/database"
	"auth/routes"
	"auth/utils"
	"auth/utils/config"
	"auth/utils/env"
	"crypto/tls"
	"embed"
	"flag"
	"net/http"
	"os"

	"codeberg.org/coldwire/cwauth"
	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
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

	// Connet to hydra public oauth provider
	cwauth.InitOauth2(oauth2.Config{
		ClientID:     "auth",
		ClientSecret: config.Conf.Server.AuthSecret,
		RedirectURL:  config.Conf.Server.AuthUrl + "/api/callback",
	}, config.Conf.Hydra.Public)
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

	// Load view as static website
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(views),
		PathPrefix: "views/public",
		Browse:     true,
	}))

	utils.InitClients()

	log.Info().Err(app.Listen(config.Conf.Server.Address + ":" + config.Conf.Server.Port))
}
