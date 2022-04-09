package utils

import (
	"auth/utils/config"

	"codeberg.org/coldwire/cwhydra"
	"github.com/rs/zerolog/log"
)

func InitClients() {
	for _, cv := range config.Conf.Hydra.Clients {
		c, _ := cwhydra.ClientManager(*cwhydra.AdminApi).Get(cv.ClientId)
		if c.ClientId != "" {
			continue
		} else {
			_, err := cwhydra.ClientManager(*cwhydra.AdminApi).Create(cv)
			if err != nil {
				log.Err(err).Msg(err.Error())
			}
		}
	}
}
