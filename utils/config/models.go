package config

import "codeberg.org/coldwire/cwhydra"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Hydra    HydraConfig
}

type ServerConfig struct {
	Address    string `toml:"address"`     // Address for webserver to listen on
	Port       string `toml:"port"`        // Port for webserver to listen on
	AuthUrl    string `toml:"auth-url"`    // Url to acces the server (https://auth.coldwire.org/)
	AuthSecret string `toml:"auth-secret"` // Oauth secret for the auth client
}

type DatabaseConfig struct {
	Driver   string                 `toml:"driver"` // postgres or sqlite
	Postgres DatabasePostgresConfig // Postgres config
	Sqlite   DatabaseSqliteConfig   // Sqlite config
}

type DatabasePostgresConfig struct {
	Address  string `toml:"address"`  // Address of the postgres instance
	Port     string `toml:"port"`     // Port of the postgres instance
	User     string `toml:"user"`     // USer of the database
	Password string `toml:"password"` // Password of the database
	Name     string `toml:"name"`     // Name of the database
}

type DatabaseSqliteConfig struct {
	Path string `toml:"path"` // path to the database file
}

type HydraConfig struct {
	Admin   string `toml:"admin"`
	Public  string `toml:"public"`
	Clients []cwhydra.OAuth2Client
}
