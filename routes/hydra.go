package routes

import (
	"auth/utils/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Hydra(app *fiber.App) {
	app.Get("/.well-known/*", func(c *fiber.Ctx) error {
		url := config.Conf.Hydra.Public + ".well-known/" + c.Params("*")
		log.Println(url)
		err := proxy.Do(c, url)
		if err != nil {
			return err
		}

		c.Response().Header.Del(fiber.HeaderServer)

		return nil
	})

	app.Get("/oauth2/*", func(c *fiber.Ctx) error {
		url := config.Conf.Hydra.Public + "oauth2/" + c.Params("*")
		err := proxy.Do(c, url)
		if err != nil {
			return err
		}

		c.Response().Header.Del(fiber.HeaderServer)

		return nil
	})
}
