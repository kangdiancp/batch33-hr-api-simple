package repositories

import (
	"context"

	"github.com/codeid/hr-api-simple/internal/models"
	"gorm.io/gorm"
)

// 1. create interface & methods
type RegionRepository interface {
	FindAll(ctx context.Context) ([]models.Region, error)
	FindByID(ctx context.Context, id uint) (*models.Region, error)
	Create(ctx context.Context, region *models.Region) error
	Update(ctx context.Context, region *models.Region) error
	Delete(ctx context.Context, id uint) error
}

// 2. RegionRepository struct
type regionRepository struct {
	*BaseRepository
}

// 2. constructor
func NewRegionRepository(db *gorm.DB) RegionRepository {
	return &regionRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// FindAll implements RegionRepository.
func (r *regionRepository) FindAll(ctx context.Context) ([]models.Region, error) {
	var regions []models.Region
	err := r.DB.WithContext(ctx).Find(&regions).Error
	return regions, err
}

func (r *regionRepository) FindByID(ctx context.Context, id uint) (*models.Region, error) {
	var region models.Region
	err := r.DB.WithContext(ctx).First(&region, id).Error
	if err != nil {
		return nil, err
	}
	return &region, nil
}

func (r *regionRepository) Create(ctx context.Context, region *models.Region) error {
	return r.DB.WithContext(ctx).Create(region).Error
}

func (r *regionRepository) Update(ctx context.Context, region *models.Region) error {
	return r.DB.WithContext(ctx).Save(region).Error
}

func (r *regionRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Region{}, id).Error
}
