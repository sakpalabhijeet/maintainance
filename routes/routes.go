package routes

import (
	"Maintainance/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	flatHandler *handlers.FlatHandler,
	societyHandler *handlers.SocietyHandler,
	ownerHandler *handlers.OwnerHandler,
	userHandler *handlers.UserHandler,
	billHandler *handlers.BillHandler,
	billItemHandler *handlers.BillItemHandler,
	maintenanceRateHandler *handlers.MaintenanceRateHandler,
) {

	api := router.Group("/api")

	// -------------------------
	// Society Routes
	// -------------------------
	api.POST("/societies", societyHandler.CreateSociety)
	// api.GET("/societies", societyHandler.GetAll)
	// api.GET("/societies/:id", societyHandler.GetByID)

	// -------------------------
	// User Routes
	// -------------------------
	api.POST("/users", userHandler.Create)
	api.GET("/users", userHandler.GetAll)
	api.GET("/users/:id", userHandler.GetByID)
	api.DELETE("/users/:id", userHandler.Delete)

	// -------------------------
	// Owner Routes
	// -------------------------
	api.POST("/owners", ownerHandler.Create)
	api.GET("/owners", ownerHandler.GetAll)
	api.GET("/owners/:id", ownerHandler.GetByID)
	api.PUT("/owners/:id", ownerHandler.Update)
	api.DELETE("/owners/:id", ownerHandler.Delete)

	// -------------------------
	// Flat Routes
	// -------------------------
	api.POST("/flats", flatHandler.Create)
	// api.GET("/flats", flatHandler.GetAll)
	api.GET("/flats/:id", flatHandler.GetByID)
	api.PUT("/flats/:id", flatHandler.Update)
	api.DELETE("/flats/:id", flatHandler.Delete)

	// -------------------------
	// Maintenance Rate Routes
	// -------------------------
	api.POST("/maintenance-rates", maintenanceRateHandler.Create)
	api.GET("/maintenance-rates", maintenanceRateHandler.GetAll)
	api.GET("/maintenance-rates/:id", maintenanceRateHandler.GetByID)
	api.PUT("/maintenance-rates/:id", maintenanceRateHandler.Update)
	api.DELETE("/maintenance-rates/:id", maintenanceRateHandler.Delete)

	// -------------------------
	// Bill Routes
	// -------------------------
	api.POST("/bills/generate", billHandler.GenerateMonthlyBills)
	// api.GET("/bills/:id", billHandler.GetByID)
	// api.GET("/bills", billHandler.GetAll)

	// -------------------------
	// Bill Item Routes
	// -------------------------
	api.GET("/billitems/:bill_id", billItemHandler.GetByBillID)
}
