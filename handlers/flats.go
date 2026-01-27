package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlatHandler struct {
	services services.FlatService
}

func NewFlatHandler( services  services.FlatService) *FlatHandler{
	return &FlatHandler{services: services}
}

func (h* FlatHandler)CreateFlat(c *gin.Context){
	var flat models.Flat

	if err:= c.ShouldBindJSON(&flat); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	if err:= h.services.CreateFlat(c.Request.Context(), &flat);err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":201,
		"success":true,
		"message":"flat created successfully!",

	})
}
