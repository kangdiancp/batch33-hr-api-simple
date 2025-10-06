package services

import (
	"context"
	"errors"

	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/repositories"
)

// 1. create service
type CountryService interface {
	GetAllCountries(ctx context.Context) ([]models.Country, error)
	GetCountryByID(ctx context.Context, id string) (*models.Country, error)
	GetCountriesByRegion(ctx context.Context, regionID uint) ([]models.Country, error)
	CreateCountry(ctx context.Context, country *models.Country) error
	UpdateCountry(ctx context.Context, country *models.Country) error
	DeleteCountry(ctx context.Context, id string) error
}

type countryService struct {
	countryRepo repositories.CountryRepository
	regionRepo  repositories.RegionRepository
}

func NewCountryService(
	countryRepo repositories.CountryRepository,
	regionRepo repositories.RegionRepository,
) CountryService {
	return &countryService{
		countryRepo: countryRepo,
		regionRepo:  regionRepo,
	}
}

func (s *countryService) GetAllCountries(ctx context.Context) ([]models.Country, error) {
	return s.countryRepo.FindAll(ctx)
}

func (s *countryService) GetCountryByID(ctx context.Context, id string) (*models.Country, error) {
	if id == "" {
		return nil, errors.New("country ID cannot be empty")
	}
	if len(id) != 2 {
		return nil, errors.New("country ID must be 2 characters")
	}
	return s.countryRepo.FindByID(ctx, id)
}

func (s *countryService) GetCountriesByRegion(ctx context.Context, regionID uint) ([]models.Country, error) {
	if regionID == 0 {
		return nil, errors.New("region ID cannot be empty")
	}
	return s.countryRepo.FindByRegionID(ctx, regionID)
}

func (s *countryService) CreateCountry(ctx context.Context, country *models.Country) error {
	if country.CountryID == "" {
		return errors.New("country ID cannot be empty")
	}
	if len(country.CountryID) != 2 {
		return errors.New("country ID must be 2 characters")
	}
	if country.CountryName == "" {
		return errors.New("country name cannot be empty")
	}
	if len(country.CountryName) > 40 {
		return errors.New("country name cannot exceed 40 characters")
	}
	if country.RegionID == 0 {
		return errors.New("region ID cannot be empty")
	}

	// Check if region exists
	_, err := s.regionRepo.FindByID(ctx, country.RegionID)
	if err != nil {
		return errors.New("region not found")
	}

	return s.countryRepo.Create(ctx, country)
}

func (s *countryService) UpdateCountry(ctx context.Context, country *models.Country) error {
	if country.CountryID == "" {
		return errors.New("country ID cannot be empty")
	}
	if country.CountryName == "" {
		return errors.New("country name cannot be empty")
	}
	if len(country.CountryName) > 40 {
		return errors.New("country name cannot exceed 40 characters")
	}
	if country.RegionID == 0 {
		return errors.New("region ID cannot be empty")
	}

	// Check if region exists
	_, err := s.regionRepo.FindByID(ctx, country.RegionID)
	if err != nil {
		return errors.New("region not found")
	}

	return s.countryRepo.Update(ctx, country)
}

func (s *countryService) DeleteCountry(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("country ID cannot be empty")
	}
	if len(id) != 2 {
		return errors.New("country ID must be 2 characters")
	}
	return s.countryRepo.Delete(ctx, id)
}
