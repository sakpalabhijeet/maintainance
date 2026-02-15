package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type BillItemHandler struct {
	service services.BillItemService
}

func NewBillItemHandler(service services.BillItemService,) *BillItemHandler {
	return &BillItemHandler{service: service}
}

func (h *BillItemHandler) Create(c *gin.Context) {
	var item models.BillItem

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.Create(c.Request.Context(), &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    item,
	})
}

func (h *BillItemHandler) GetByBillID(c *gin.Context) {
	billID, err := strconv.ParseInt(c.Param("billId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid bill id",
		})
		return
	}

	items, err := h.service.GetByBillID(c.Request.Context(), billID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
	})
}

func (h *BillItemHandler) DeleteByBillID(c *gin.Context) {
	billID, err := strconv.ParseInt(c.Param("billId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid bill id",
		})
		return
	}

	if err := h.service.DeleteByBillID(c.Request.Context(), billID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "bill items deleted successfully",
	})
}
