package main

import (
	"Maintainance/database"
	"Maintainance/handlers"
	"Maintainance/internal/config"
	"Maintainance/repo"
	"Maintainance/routes"
	"Maintainance/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg:= config.Load()
	db:= database.Connect(cfg)

	societyRepo := repo.NewSocietyRepository(db)
	societyService:= services.NewSocietyService(societyRepo)
	societyHandler:= handlers.NewSocietyHandler(societyService)
	// -----------------------------------------------------
	flatRepo := repo.NewFlatRepository(db)
	flatService:= services.NewFlatService(flatRepo)
	flatHandler:= handlers.NewFlatHandler(flatService)
	// -----------------------------------------------------

	router := gin.Default()
	api:= router.Group("/api")
	routes.RegisterFlatRoutes(api,flatHandler)
	routes.RegisterSocietyRoutes(api,societyHandler)
	router.Run(":8080")
}