package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	countryService services.CountryService
}

func NewCountryHandler(countryService services.CountryService) *CountryHandler {
	return &CountryHandler{
		countryService: countryService,
	}
}

func (h *CountryHandler) GetCountries(c *gin.Context) {
	countries, err := h.countryService.GetAllCountries(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to fetch countries",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    countries,
	})
}

func (h *CountryHandler) GetCountry(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Country ID must be 2 characters",
		})
		return
	}

	country, err := h.countryService.GetCountryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "Country not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    country,
	})
}

func (h *CountryHandler) GetCountriesByRegion(c *gin.Context) {
	regionIDStr := c.Param("region_id")
	regionID, err := strconv.ParseUint(regionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}

	countries, err := h.countryService.GetCountriesByRegion(c.Request.Context(), uint(regionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to fetch countries",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    countries,
	})
}

func (h *CountryHandler) CreateCountry(c *gin.Context) {
	var country models.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}

	if err := h.countryService.CreateCountry(c.Request.Context(), &country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create country",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Country created successfully",
		"data":    country,
	})
}

func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Country ID must be 2 characters",
		})
		return
	}

	var country models.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}

	country.CountryID = id

	if err := h.countryService.UpdateCountry(c.Request.Context(), &country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to update country",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Country updated successfully",
		"data":    country,
	})
}

func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Country ID must be 2 characters",
		})
		return
	}

	if err := h.countryService.DeleteCountry(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to delete country",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Country deleted successfully",
	})
}
