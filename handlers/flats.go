package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type FlatHandler struct {
	services services.FlatService
}

func NewFlatHandler( services  services.FlatService) *FlatHandler{
	return &FlatHandler{services: services}
}
func (h *FlatHandler) Create(c *gin.Context) {
	var flat models.Flat

	if err := c.ShouldBindJSON(&flat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.services.Create(c.Request.Context(), &flat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    flat,
	})
}

func (h *FlatHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid id",
		})
		return
	}

	flat, err := h.services.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "flat not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    flat,
	})
}


func (h *FlatHandler) Update(c *gin.Context) {
	var flat models.Flat

	if err := c.ShouldBindJSON(&flat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if flat.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "id is required for update",
		})
		return
	}

	if err := h.services.Update(c.Request.Context(), &flat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "flat updated successfully",
	})
}

func (h *FlatHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid id",
		})
		return
	}

	if err := h.services.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "flat deleted successfully",
	})
}


