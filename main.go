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
	ownerRepo := repo.NewOwnerRepository(db)
	ownerService:= services.NewOwnerService(ownerRepo)
	ownerHandler:= handlers.NewOwnerHandler(ownerService)
	// -----------------------------------------------------
	userRepo:= repo.NewUserRepository(db)
	userService:= services.NewUserService(userRepo)
	userHandler:= handlers.NewUserHandler(userService)
	// -----------------------------------------------------
	billRepo:= repo.NewBillRepository(db)
	billService:= services.NewBillService(billRepo)
	billHandler:= handlers.NewBillHandler(billService)
	// -----------------------------------------------------
	router := gin.Default()
	api:= router.Group("/api")
	routes.RegisterFlatRoutes(api,flatHandler)
	routes.RegisterSocietyRoutes(api,societyHandler)
	routes.RegisterOwnerRoutes(api,ownerHandler)
	routes.RegisterUserRoutes(api, userHandler)
	routes.RegisterBillRoutes(api,*billHandler)
	router.Run(":8080")
}