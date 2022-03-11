package database

import (
	"auth/utils/config"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Expose database to the whole project
var DB *gorm.DB
var Ready bool = false

// Connect to the database
func Connect() error {
	var err error
	var driver gorm.Dialector

	newLogger := logger.New(
		&log.Logger,
		logger.Config{
			LogLevel: logger.Silent,
		},
	)

	var gconf *gorm.Config = &gorm.Config{
		Logger: newLogger,
	}

	log.Info().Msg("Using driver: " + config.Conf.Database.Driver)

	switch config.Conf.Database.Driver {
	case "sqlite":
		if config.Conf.Database.Sqlite.Path != "" {
			driver = sqlite.Open(config.Conf.Database.Sqlite.Path)
		} else {
			log.Fatal().Msg("Please set a path to the database file or use Postgres if you want a remote database")
		}
	case "postgres":
		DSN := []string{
			"postgresql://",
			config.Conf.Database.Postgres.Address,
			":" + config.Conf.Database.Postgres.Port,
			"/" + config.Conf.Database.Postgres.Name,
			"?user=" + config.Conf.Database.Postgres.User,
			"&password=" + config.Conf.Database.Postgres.Password,
		}

		driver = postgres.Open(strings.Join(DSN[:], ""))
	default:
		log.Fatal().Msg("Database driver not found!")
	}

	// Run until the database is connected
	for {
		DB, err = gorm.Open(driver, gconf)

		if err != nil {
			Ready = false
			log.Error().Msg(err.Error())
			time.Sleep(5 * time.Second)
		} else {
			log.Info().Msg("Successfully connected to the database!")
			Ready = true
			break
		}
	}

	return nil
}
