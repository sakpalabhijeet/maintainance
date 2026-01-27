package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BillHandler struct {
	services services.BillService
}

func NewBillHandler ( services services.BillService) *BillHandler{
	return &BillHandler{
		services: services,
	}
}

func (h *BillHandler)CreateBill( c *gin.Context){
	var bill models.Bill

	if err:= c.ShouldBindJSON(&bill);err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	if err:= h.services.CreateBill(c.Request.Context(), &bill);err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": 201,
		"success": true,
		"message": "bill created successfully!",
	})
}