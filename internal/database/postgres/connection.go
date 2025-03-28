package database

import (
	"os"

	"github.com/Afthaab/Sales-Report-Lumel/internal/model/dbmodel"
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

	_, err = db.DB()
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to get the database instance")
	}

	err = db.AutoMigrate(&dbmodel.Customer{}, &dbmodel.Region{}, &dbmodel.Category{}, &dbmodel.Product{}, &dbmodel.Order{}, &dbmodel.OrderItem{})
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to migrate the tables")
	}

	return db
}
