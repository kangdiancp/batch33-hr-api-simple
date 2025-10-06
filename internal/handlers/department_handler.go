package handlers

import (
	"net/http"

	"github.com/codeid/hr-api-simple/internal/services"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	deparmentService services.DepartmentService
}

func NewDeparmentHandler(departmentService services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		deparmentService: departmentService,
	}
}

func (d *DepartmentHandler) GetDepartments(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "{}",
	})
}
