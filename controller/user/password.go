package user

import (
	"auth/models"
	"auth/utils/errors"
	"auth/utils/tokens"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func Password(c *fiber.Ctx) error {
	token := c.Cookies("token")

	request := struct {
		Password string `json:"password"`
		New      string `json:"new"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		return errors.Handle(c, errors.ErrBody, err)
	}

	t, err := tokens.Parse(token)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	user := models.User{
		Name: t.Username,
	}

	user, err = user.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if user.Name != t.Username {
		return errors.Handle(c, errors.ErrPermission)
	}

	valid, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	if !valid {
		return errors.Handle(c, errors.ErrAuthPassword)
	}

	newPassword, err := argon2id.CreateHash(request.New, argon2id.DefaultParams)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	err = user.SetPassword(newPassword)
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseUpdate, err)
	}

	return errors.Handle(c, errors.Success)
}
