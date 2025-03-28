package server

import (
	"github.com/Afthaab/Sales-Report-Lumel/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func StartApplication(handler handler.HandlerInterface) {
	router := gin.Default()

	registerRoutes(router, handler)

	log.Info().Msg("Server is running on http://localhost:8000")

	err := router.Run(":8000")
	if err != nil {
		log.Fatal().AnErr("failed to start the server", err)
	}
}
