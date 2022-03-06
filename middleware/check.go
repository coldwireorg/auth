package middleware

import (
	"auth/oauth"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CheckAuthenticated(c *fiber.Ctx) error {
	idToken := c.Cookies("id_token")
	accesToken := c.Cookies("access_token")
	challenge := c.Query("login_challenge")

	csrf := c.Cookies("oauth2_authentication_csrf")
	csrfInsecure := c.Cookies("oauth2_authentication_csrf_insecure")

	isTokenValide := oauth.CheckToken(idToken, accesToken)
	if isTokenValide {
		return c.Redirect("/user")
	}

	if challenge != "" && (csrf != "" || csrfInsecure != "") {
		return c.Next()
	}

	url, err := oauth.AuthURL()
	if err != nil {
		log.Println(err)
	}

	return c.Redirect(url)
}

func CheckUser(c *fiber.Ctx) error {
	idToken := c.Cookies("id_token")
	accesToken := c.Cookies("access_token")

	isTokenValide := oauth.CheckToken(idToken, accesToken)
	log.Print(isTokenValide)
	if !isTokenValide {
		url, err := oauth.AuthURL()
		if err != nil {
			log.Println(err)
		}

		return c.Redirect(url)
	}

	return c.Next()
}
