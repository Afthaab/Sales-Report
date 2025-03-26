package main

import (
	"fmt"

	database "github.com/Afthaab/Sales-Report-Lumel/internal/database/postgres"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("faile to load the env file")
	}
}

func main() {
	db := database.ConnectToDatabase()
	fmt.Println(db)
}
