package services

import (
	"context"

	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/repositories"
)

type DepartmentService interface {
	GetAllDepartment(ctx context.Context) ([]models.Department, error)
	GetDepartmentByID(ctx context.Context, id string) (*models.Department, error)
	CreateDepartment(ctx context.Context, Department *models.Department) error
	UpdateDepartment(ctx context.Context, Department *models.Department) error
	DeleteDepartment(ctx context.Context, id string) error
}

// 2 department struct
type departmentService struct {
	departmentRepo repositories.DepartmentRepository
}

// 3. constructor
func NewDepartmentService(departmentRepo repositories.DepartmentRepository) DepartmentService {
	return &departmentService{
		departmentRepo: departmentRepo,
	}
}

// CreateDepartment implements DepartmentService.
func (d *departmentService) CreateDepartment(ctx context.Context, Department *models.Department) error {
	panic("unimplemented")
}

// DeleteDepartment implements DepartmentService.
func (d *departmentService) DeleteDepartment(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetAllDepartment implements DepartmentService.
func (d *departmentService) GetAllDepartment(ctx context.Context) ([]models.Department, error) {
	panic("unimplemented")
}

// GetDepartmentByID implements DepartmentService.
func (d *departmentService) GetDepartmentByID(ctx context.Context, id string) (*models.Department, error) {
	panic("unimplemented")
}

// UpdateDepartment implements DepartmentService.
func (d *departmentService) UpdateDepartment(ctx context.Context, Department *models.Department) error {
	panic("unimplemented")
}
