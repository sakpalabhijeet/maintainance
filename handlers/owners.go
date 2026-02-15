package handlers

import (
	"Maintainance/models"
	"Maintainance/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OwnerHandler struct {
	service services.OwnerService
}

func NewOwnerHandler(service services.OwnerService) *OwnerHandler {
	return &OwnerHandler{service: service}
}

// POST /api/owners
func (h *OwnerHandler) Create(c *gin.Context) {
	var owner models.Owners

	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.Create(c.Request.Context(), &owner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    owner,
	})
}

// GET /api/owners/:id
func (h *OwnerHandler) GetByID(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid owner id",
		})
		return
	}

	owner, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "owner not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owner,
	})
}

// GET /api/owners
func (h *OwnerHandler) GetAll(c *gin.Context) {

	owners, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owners,
	})
}

// PUT /api/owners/:id
func (h *OwnerHandler) Update(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid owner id",
		})
		return
	}

	var owner models.Owners
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	owner.Id = id

	if err := h.service.Update(c.Request.Context(), &owner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owner,
	})
}

// DELETE /api/owners/:id
func (h *OwnerHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid owner id",
		})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "owner deleted successfully",
	})
}
