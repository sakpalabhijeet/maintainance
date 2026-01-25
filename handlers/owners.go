package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OwnerHandler struct {
	services services.OwnerService
}

func NewOwnerHandler (service services.OwnerService) *OwnerHandler{
	return &OwnerHandler{services: service}
}

func (h* OwnerHandler)CreateOwner(c *gin.Context){
	var owner models.Owners

	if err:= c.ShouldBindJSON(&owner);err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	if err:= h.services.CreateOwner(c.Request.Context(), &owner);err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":201,
		"success":true,
		"message": "owner created successfully!",
	})
} 