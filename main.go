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

	// ---------------------------
	// Load Config & Connect DB
	// ---------------------------
	cfg := config.Load()
	db := database.Connect(cfg)

	// ---------------------------
	// Repositories
	// ---------------------------
	societyRepo := repo.NewSocietyRepository(db)
	flatRepo := repo.NewFlatRepository(db)
	ownerRepo := repo.NewOwnerRepository(db)
	userRepo := repo.NewUserRepository(db)
	billRepo := repo.NewBillRepository(db)
	billItemRepo := repo.NewBillItemRepository(db)
	maintenanceRateRepo := repo.NewMaintenanceRateRepository(db)

	// ---------------------------
	// Services
	// ---------------------------
	societyService := services.NewSocietyService(societyRepo)
	flatService := services.NewFlatService(flatRepo)
	ownerService := services.NewOwnerService(ownerRepo)
	userService := services.NewUserService(userRepo)

	billService := services.NewBillService(
		billRepo,
		flatRepo,
		maintenanceRateRepo,
		billItemRepo,
	)

	maintenanceRateService := services.NewMaintenanceRateService(maintenanceRateRepo)
	billItemService := services.NewBillItemService(billItemRepo)

	// ---------------------------
	// Handlers
	// ---------------------------
	societyHandler := handlers.NewSocietyHandler(societyService)
	flatHandler := handlers.NewFlatHandler(flatService)
	ownerHandler := handlers.NewOwnerHandler(ownerService)
	userHandler := handlers.NewUserHandler(userService)
	billHandler := handlers.NewBillHandler(billService)
	maintenanceRateHandler := handlers.NewMaintenanceRateHandler(maintenanceRateService)
	billItemHandler := handlers.NewBillItemHandler(billItemService)

	// ---------------------------
	// Router
	// ---------------------------
	router := gin.Default()

	routes.RegisterRoutes(
		router,
		flatHandler,
		societyHandler,
		ownerHandler,
		userHandler,
		billHandler,
		billItemHandler,
		maintenanceRateHandler,
	)

	// ---------------------------
	// Start Server
	// ---------------------------
	router.Run(":8080")
}
