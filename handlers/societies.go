package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocietyHandler struct {
	services services.SocietyService
}

func NewSocietyHandler( services services.SocietyService)*SocietyHandler{
	return &SocietyHandler{services: services}
}

func (h *SocietyHandler)CreateSociety(c *gin.Context){
	var society models.Society

	if err:= c.ShouldBindJSON(&society);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	if err:= h.services.CreateSociety(c.Request.Context(),&society);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": 201,
		"success": true,
		"message":"Society created successfully!",
	})
}