package repositories

import (
	"context"

	"github.com/codeid/hr-api-simple/internal/models"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindAll(ctx context.Context) ([]models.Department, error)
	FindByID(ctx context.Context, id string) (*models.Department, error)
	FindByRegionID(ctx context.Context, regionID uint) ([]models.Department, error)
	Create(ctx context.Context, Department *models.Department) error
	Update(ctx context.Context, Department *models.Department) error
	Delete(ctx context.Context, id string) error
}

type departmentRepository struct {
	*BaseRepository
}

// 2.constructor
func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// Create implements DepartmentRepository.
func (d *departmentRepository) Create(ctx context.Context, Department *models.Department) error {
	panic("unimplemented")
}

// Delete implements DepartmentRepository.
func (d *departmentRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements DepartmentRepository.
func (d *departmentRepository) FindAll(ctx context.Context) ([]models.Department, error) {
	panic("unimplemented")
}

// FindByID implements DepartmentRepository.
func (d *departmentRepository) FindByID(ctx context.Context, id string) (*models.Department, error) {
	panic("unimplemented")
}

// FindByRegionID implements DepartmentRepository.
func (d *departmentRepository) FindByRegionID(ctx context.Context, regionID uint) ([]models.Department, error) {
	panic("unimplemented")
}

// Update implements DepartmentRepository.
func (d *departmentRepository) Update(ctx context.Context, Department *models.Department) error {
	panic("unimplemented")
}
