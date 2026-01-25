package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	services services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{services: service}
}

func (h *UserHandler)CreateUser (c *gin.Context){
	var user models.Users

	if err:= c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	if err:= h.services.CreateUser(c.Request.Context(), &user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":201,
		"success":true,
		"message":"user created successfully!",
	})
}