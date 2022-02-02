package models

import (
	"auth/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type User struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	PublicKey  []byte `db:"public_key"`
	PrivateKey []byte `db:"private_key"`
}

func (user User) Exist() bool {
	usr, err := user.Get()
	if err == pgx.ErrNoRows {
		return false
	}

	if usr.Username == user.Username {
		return true
	}

	return false
}

func (user User) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO users(username, password, public_key, private_key) VALUES($1, $2, $3, $4)`, user.Username, user.Password, user.PublicKey, user.PrivateKey)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (user User) Get() (User, error) {
	err := pgxscan.Get(context.Background(), database.DB, &user, `SELECT * FROM users WHERE username = $1`, user.Username)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}

	return user, nil
}

func (user User) Pubkey() ([]byte, error) {
	var pubkey []byte
	err := pgxscan.Get(context.Background(), database.DB, &pubkey, `SELECT public_key FROM users WHERE username = $1`, user.Username)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return pubkey, nil
}
