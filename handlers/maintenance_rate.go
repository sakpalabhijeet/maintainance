package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MaintenanceRateHandler struct {
	service services.MaintenanceRateService
}

func NewMaintenanceRateHandler(service services.MaintenanceRateService) *MaintenanceRateHandler {
	return &MaintenanceRateHandler{service: service}
}

func (h *MaintenanceRateHandler) Create(c *gin.Context) {

	var rate models.MaintenanceRate

	if err := c.ShouldBindJSON(&rate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if rate.StartDate.IsZero() {
		rate.StartDate = time.Now()
	}

	if err := h.service.Create(c.Request.Context(), &rate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    rate,
	})
}

func (h *MaintenanceRateHandler) GetAll(c *gin.Context) {

	rates, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rates,
	})
}

func (h *MaintenanceRateHandler) GetByID(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid id",
		})
		return
	}

	rate, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "rate not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rate,
	})
}

func (h *MaintenanceRateHandler) Update(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid id"})
		return
	}

	var rate models.MaintenanceRate
	if err := c.ShouldBindJSON(&rate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	rate.ID = id

	if err := h.service.Update(c.Request.Context(), &rate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rate})
}

func (h *MaintenanceRateHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid id"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "rate deleted successfully"})
}
