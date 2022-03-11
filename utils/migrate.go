package utils

import (
	"auth/database"
	"auth/models"

	"github.com/rs/zerolog/log"
)

// Migrate database tables
func MigrateTables() {
	//database.DB.Migrator().DropTable(&models.User{})
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
