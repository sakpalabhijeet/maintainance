package routes

import (
	"Maintainance/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterFlatRoutes(r *gin.RouterGroup, handler * handlers.FlatHandler){
	r.POST("/flats", handler.CreateFlat)
}