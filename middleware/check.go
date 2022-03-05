package middleware

import (
	"auth/utils"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func Check(c *fiber.Ctx) error {
	challenge := c.Query("login_challenge")

	if challenge != "" {
		return c.Next()
	}

	state, err := utils.GenerateRandomString(16)
	if err != nil {
		return c.Next()
	}

	params := url.Values{
		"client_id":     []string{"auth"},
		"max_age":       []string{"0"},
		"redirect_uri":  []string{"http://127.0.0.1:3002/user"},
		"response_type": []string{"code"},
		"scope": []string{
			"openid",
			"offline",
		},
		"state": []string{state},
	}

	return c.Redirect("http://127.0.0.1:4444/oauth2/auth?" + params.Encode())
}
