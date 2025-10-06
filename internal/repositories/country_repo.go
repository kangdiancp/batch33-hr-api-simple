package repositories

import (
	"context"

	"github.com/codeid/hr-api-simple/internal/models"
	"gorm.io/gorm"
)

// 1. create interface & methods
type CountryRepository interface {
	FindAll(ctx context.Context) ([]models.Country, error)
	FindByID(ctx context.Context, id string) (*models.Country, error)
	FindByRegionID(ctx context.Context, regionID uint) ([]models.Country, error)
	Create(ctx context.Context, country *models.Country) error
	Update(ctx context.Context, country *models.Country) error
	Delete(ctx context.Context, id string) error
}

type countryRepository struct {
	*BaseRepository
}

func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *countryRepository) FindAll(ctx context.Context) ([]models.Country, error) {
	var countries []models.Country
	err := r.DB.WithContext(ctx).Preload("Region").Find(&countries).Error
	return countries, err
}

func (r *countryRepository) FindByID(ctx context.Context, id string) (*models.Country, error) {
	var country models.Country
	err := r.DB.WithContext(ctx).Preload("Region").First(&country, "country_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) FindByRegionID(ctx context.Context, regionID uint) ([]models.Country, error) {
	var countries []models.Country
	err := r.DB.WithContext(ctx).Preload("Region").Where("region_id = ?", regionID).Find(&countries).Error
	return countries, err
}

func (r *countryRepository) Create(ctx context.Context, country *models.Country) error {
	return r.DB.WithContext(ctx).Create(country).Error
}

func (r *countryRepository) Update(ctx context.Context, country *models.Country) error {
	return r.DB.WithContext(ctx).Save(country).Error
}

func (r *countryRepository) Delete(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Delete(&models.Country{}, "country_id = ?", id).Error
}
