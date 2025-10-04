package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/services"
	"github.com/gin-gonic/gin"
)

type RegionHandler struct {
	regionService services.RegionService
}

func NewRegionHandler(regionService services.RegionService) *RegionHandler {
	return &RegionHandler{
		regionService: regionService,
	}
}

func (h *RegionHandler) GetRegions(c *gin.Context) {
	regions, err := h.regionService.GetAllRegions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to fetch regions",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    regions,
	})
}

func (h *RegionHandler) GetRegion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}

	region, err := h.regionService.GetRegionByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "Region not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    region,
	})
}

func (h *RegionHandler) CreateRegion(c *gin.Context) {
	var region models.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}

	if err := h.regionService.CreateRegion(c.Request.Context(), &region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create region",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Region created successfully",
		"data":    region,
	})
}

func (h *RegionHandler) UpdateRegion(c *gin.Context) {
	idStr := c.Param("id")
	//convert from string to uint, base10, format 32bit
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}

	var region models.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}

	region.RegionID = uint(id)

	if err := h.regionService.UpdateRegion(c.Request.Context(), &region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to update region",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Region updated successfully",
		"data":    region,
	})
}

func (h *RegionHandler) DeleteRegion(c *gin.Context) {
	idStr := c.Param("id")
	//convert from string to uint, base10, format 32bit
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}

	if err := h.regionService.DeleteRegion(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to delete region",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Region deleted successfully",
	})
}
