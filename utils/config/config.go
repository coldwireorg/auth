package config

import (
	"auth/utils/env"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Export config to the whole project
var Conf *Config

// Init config to ensure that env variable are loaded
func Init(file string) {
	// if config file is specified we use it
	if file != "" {
		log.Info().Msg("Loading config file: " + file)
		f, err := os.ReadFile(file)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		var c Config
		_, err = toml.Decode(string(f), &c)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		Conf = &c

		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg(err.Error())
	}

	Conf = &Config{
		Server: ServerConfig{
			Address: env.Get("AUTH_SERVER_ADDRESS", "0.0.0.0"),
			Port:    env.Get("AUTH_SERVER_PORT", "3000"),
		},

		Database: DatabaseConfig{
			Driver: env.Get("AUTH_DATABASE_DRIVER", "sqlite"),
			Postgres: DatabasePostgresConfig{
				Address:  env.Get("AUTH_DATABASE_ADDRESS", "127.0.0.1"),
				Port:     env.Get("AUTH_DATABASE_PORT", "5432"),
				User:     env.Get("AUTH_DATABASE_USER", "postgres"),
				Password: env.Get("AUTH_DATABASE_PASSWORD", "123456789"),
				Name:     env.Get("AUTH_DATABASE_NAME", "AUTH"),
			},
			Sqlite: DatabaseSqliteConfig{
				Path: env.Get("AUTH_DATABASE_PATH", "/tmp/AUTH.sqlite"),
			},
		},

		Hydra: HydraConfig{
			Public: env.Get("AUTH_HYDRA_PUBLIC_URL", "http://127.0.0.1:4444/"),
			Admin:  env.Get("AUTH_HYDRA_ADMIN_URL", "http://127.0.0.1:4445"),
		},

		Oauth: OauthConfig{
			Server:   env.Get("AUTH_OAUTH_SERVER", ""),
			Id:       env.Get("AUTH_OAUTH_CLIENT", "AUTH"),
			Secret:   env.Get("AUTH_OAUTH_SECRET", ""),
			Callback: env.Get("AUTH_OAUTH_CALLBACK", "https://AUTH.coldwire.org/api/user/auth/oauth2/callback"),
		},
	}
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Hydra    HydraConfig
	Oauth    OauthConfig
}

type ServerConfig struct {
	Address string // Address for webserver to listen on
	Port    string // Port for webserver to listen on
}

type DatabaseConfig struct {
	Driver   string                 // postgres or sqlite
	Postgres DatabasePostgresConfig // Postgres config
	Sqlite   DatabaseSqliteConfig   // Sqlite config
}

type DatabasePostgresConfig struct {
	Address  string // Address of the postgres instance
	Port     string // Port of the postgres instance
	User     string // USer of the database
	Password string // Password of the database
	Name     string // Name of the database
}

type DatabaseSqliteConfig struct {
	Path string // path to the database file
}

type HydraConfig struct {
	Admin   string
	Public  string
	Clients []HydraClientsConfig
}

type HydraClientsConfig struct {
	Id            string
	GrantTypes    []string
	ResponseTypes []string
	Scope         []string
	Callback      string
	AuthMethod    string
}

type OauthConfig struct {
	Server   string
	Id       string
	Secret   string
	Callback string
}
