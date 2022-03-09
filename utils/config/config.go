package config

import (
	"auth/hydra"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Clients []hydra.OAuth2Client
}

// TODO: finish this
func Load(path string) {
	if path == "" {
		return
	}

	log.Println("Loading config file: " + path)
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	_, err = toml.Decode(string(f), &c)
	if err != nil {
		log.Fatal(err)
	}
}
