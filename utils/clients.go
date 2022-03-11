package utils

import (
	"auth/utils/config"

	"codeberg.org/coldwire/cwhydra"
	"github.com/rs/zerolog/log"
)

func InitClients() {
	clients, err := cwhydra.ClientManager(*cwhydra.AdminApi).List()
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	// If there is no client
	if len(clients) == 0 {
		_, err := cwhydra.ClientManager(*cwhydra.AdminApi).Create(cwhydra.OAuth2Client{
			ClientId:     "auth",
			ClientSecret: config.Conf.Server.AuthSecret,
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
				config.Conf.Server.AuthUrl + "/api/callback",
			},
			TokenEndpointAuthMethod: "none",
		})

		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}

	for _, cv := range config.Conf.Hydra.Clients {
		c, _ := cwhydra.ClientManager(*cwhydra.AdminApi).Get(cv.ClientId)
		if c.ClientId != "" {
			break
		}

		_, err := cwhydra.ClientManager(*cwhydra.AdminApi).Create(cv)
		if err != nil {
			log.Err(err).Msg(err.Error())
		}
	}
}
