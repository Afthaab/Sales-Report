package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("faile to load the env file")
	}
}
