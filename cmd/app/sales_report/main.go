package main

import (
	"github.com/Afthaab/Sales-Report-Lumel/internal/config"
	database "github.com/Afthaab/Sales-Report-Lumel/internal/database/postgres"
	"github.com/Afthaab/Sales-Report-Lumel/internal/handler"
	"github.com/Afthaab/Sales-Report-Lumel/internal/loader"
	"github.com/Afthaab/Sales-Report-Lumel/internal/repository"
	"github.com/Afthaab/Sales-Report-Lumel/internal/script"
	"github.com/Afthaab/Sales-Report-Lumel/internal/server"
	"github.com/Afthaab/Sales-Report-Lumel/internal/service"
)

func main() {
	config.LoadEnv()

	db := database.ConnectToDatabase()

	repo := repository.NewRepoLayer(db)

	loader := loader.NewLoader(db, repo)

	script.StartCronJob(loader) // cron job

	svc := service.NewServiceLayer(repo)

	handler := handler.NewHandlerLayer(loader, svc)

	server.StartApplication(handler)
}
