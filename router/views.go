package router

import (
	"auth/middleware"
	"context"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func SetupViews(app *fiber.App) {
	app.Get("/", middleware.Check, func(c *fiber.Ctx) error {
		return c.Redirect("/sign-in")
	})

	app.Get("/sign-in", middleware.Check, func(c *fiber.Ctx) error {
		err := c.Query("err")
		challenge := c.Query("login_challenge")

		log.Println(err)

		return c.Render("index", fiber.Map{
			"Register":  false,
			"Challenge": challenge,
			"Error":     err,
		}, "layouts/auth")
	})

	app.Get("/sign-up", middleware.Check, func(c *fiber.Ctx) error {
		err := c.Query("err")
		challenge := c.Query("login_challenge")

		return c.Render("index", fiber.Map{
			"Register":  true,
			"Challenge": challenge,
			"Error":     err,
		}, "layouts/auth")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		code := c.Query("code")

		provider, err := oidc.NewProvider(context.Background(), "http://127.0.0.1:4444/")
		if err != nil {
			log.Println(err)
		}

		config := oauth2.Config{
			ClientID:     "auth",
			ClientSecret: "secret",
			Endpoint:     provider.Endpoint(),
			RedirectURL:  "http://127.0.0.1:3002/user",
			Scopes:       []string{oidc.ScopeOpenID},
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: config.ClientID})

		if code == "" {
			return c.Redirect("/sign-in")
		}

		oauth2Token, err := config.Exchange(context.Background(), code)
		if err != nil {
			log.Println(err)
		}

		if oauth2Token == nil {
			return c.Redirect("/sign-in")
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			return c.Redirect("/sign-in")
		}

		// Parse and verify ID Token payload.
		idToken, err := verifier.Verify(context.Background(), rawIDToken)
		if err != nil {
			log.Println(err)
		}

		var out interface{}
		err = idToken.Claims(&out)
		if err != nil {
			log.Println(err)
		}

		log.Println(out)

		return c.Render("board", fiber.Map{}, "layouts/user")
	})
}
