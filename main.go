package main

import (
	"auth/database"
	"auth/middleware"
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
	"github.com/gofiber/fiber/v2/middleware/proxy"
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
	tokens.Init(env.Get("JWT_KEY", ""))
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

	if env.Get("DEV_FRONT_URL", "") == "" {
		// Load view as static website
		app.Use("/", middleware.OnAuth, filesystem.New(filesystem.Config{
			Root:       http.FS(views),
			PathPrefix: "views/public",
			Browse:     true,
		}))
	} else {
		app.Get("/*", middleware.OnAuth, func(c *fiber.Ctx) error {
			url := env.Get("DEV_FRONT_URL", "") + c.Params("*")
			err := proxy.Do(c, url)
			if err != nil {
				return err
			}

			c.Response().Header.Del(fiber.HeaderServer)

			return nil
		})
	}

	utils.InitClients()

	log.Info().Err(app.Listen(config.Conf.Server.Address + ":" + config.Conf.Server.Port))
}
