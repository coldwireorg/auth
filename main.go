package main

import (
	"auth/database"
	"auth/routes"
	"auth/utils"
	"auth/utils/config"
	"auth/utils/env"
	"crypto/tls"
	"flag"
	"net/http"
	"os"

	"codeberg.org/coldwire/cwauth"
	"codeberg.org/coldwire/cwhydra"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

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
		ClientID:     config.Conf.Oauth.Id,
		ClientSecret: config.Conf.Oauth.Secret,
		RedirectURL:  config.Conf.Oauth.Callback,
	}, config.Conf.Oauth.Server)
}

func main() {
	// Create fiber instance
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Add static server
	app.Static("/static", "static")

	// migrate database
	utils.MigrateTables()

	// Include cors
	app.Use(cors.New())

	// Setup routes
	routes.Api(app)
	routes.View(app)

	client, err := cwhydra.ClientManager(*cwhydra.AdminApi).List()
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	// If there is no client
	if len(client) == 0 {
		cwhydra.ClientManager(*cwhydra.AdminApi).Create(cwhydra.OAuth2Client{
			ClientId: "auth",
			GrantTypes: []string{
				"authorization_code",
				"refresh_token",
			},
			ResponseTypes: []string{
				"code",
				"id_token",
			},
			Scope: "openid,offline",
			RedirectUris: []string{
				"http://127.0.0.1:3002/api/callback",
			},
			TokenEndpointAuthMethod: "none",
		})
	}

	log.Info().Err(app.Listen(config.Conf.Server.Address + ":" + config.Conf.Server.Port))
}
