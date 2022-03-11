package middleware

import (
	"log"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
)

func CheckAuthenticated(c *fiber.Ctx) error {
	idToken := c.Cookies("id_token")
	accesToken := c.Cookies("access_token")
	challenge := c.Query("login_challenge")

	csrf := c.Cookies("oauth2_authentication_csrf")
	csrfInsecure := c.Cookies("oauth2_authentication_csrf_insecure")

	isTokenValide := cwauth.CheckToken(idToken, accesToken)
	if isTokenValide {
		return c.Redirect("/user")
	}

	if challenge != "" && (csrf != "" || csrfInsecure != "") {
		return c.Next()
	}

	url, err := cwauth.AuthURL()
	if err != nil {
		log.Println(err)
	}

	return c.Redirect(url)
}

func CheckUser(c *fiber.Ctx) error {
	idToken := c.Cookies("id_token")
	accesToken := c.Cookies("access_token")

	isTokenValide := cwauth.CheckToken(idToken, accesToken)
	if !isTokenValide {
		url, err := cwauth.AuthURL()
		if err != nil {
			log.Println(err)
		}

		return c.Redirect(url)
	}

	return c.Next()
}
