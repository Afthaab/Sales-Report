package database

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to connect to the database")
	} else {
		log.Info().Msg("successfully connected to the database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to get the database instance")
	}
	defer sqlDB.Close()
	return db
}
