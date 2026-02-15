package handlers

import (
	"Maintainance/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BillHandler struct {
	service services.BillService
}

func NewBillHandler(service services.BillService) *BillHandler {
	return &BillHandler{service: service}
}


type GenerateBillRequest struct {
	BillMonth string `json:"bill_month" binding:"required"` // format: 2006-01-02
	DueDate   string `json:"due_date" binding:"required"`   // format: 2006-01-02
}


func (h *BillHandler) GenerateMonthlyBills(c *gin.Context) {

	var req GenerateBillRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Parse dates (YYYY-MM-DD)
	billMonth, err := time.Parse("2006-01-02", req.BillMonth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid bill_month format (use YYYY-MM-DD)",
		})
		return
	}

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid due_date format (use YYYY-MM-DD)",
		})
		return
	}

	err = h.service.GenerateMonthlyBills(
		c.Request.Context(),
		billMonth,
		dueDate,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "monthly bills generated successfully",
	})
}
