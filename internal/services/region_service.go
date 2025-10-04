package services

import (
	"context"
	"errors"

	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/repositories"
)

// 1. create service
type RegionService interface {
	GetAllRegions(ctx context.Context) ([]models.Region, error)
	GetRegionByID(ctx context.Context, id uint) (*models.Region, error)
	CreateRegion(ctx context.Context, region *models.Region) error
	UpdateRegion(ctx context.Context, region *models.Region) error
	DeleteRegion(ctx context.Context, id uint) error
}

type regionService struct {
	regionRepo repositories.RegionRepository
}

func NewRegionService(regionRepo repositories.RegionRepository) RegionService {
	return &regionService{
		regionRepo: regionRepo,
	}
}

func (s *regionService) GetAllRegions(ctx context.Context) ([]models.Region, error) {
	return s.regionRepo.FindAll(ctx)
}

func (s *regionService) GetRegionByID(ctx context.Context, id uint) (*models.Region, error) {
	if id == 0 {
		return nil, errors.New("region ID cannot be empty")
	}
	return s.regionRepo.FindByID(ctx, id)
}

func (s *regionService) CreateRegion(ctx context.Context, region *models.Region) error {
	if region.RegionName == "" {
		return errors.New("region name cannot be empty")
	}
	if len(region.RegionName) > 25 {
		return errors.New("region name cannot exceed 25 characters")
	}
	return s.regionRepo.Create(ctx, region)
}

func (s *regionService) UpdateRegion(ctx context.Context, region *models.Region) error {
	if region.RegionID == 0 {
		return errors.New("region ID cannot be empty")
	}
	if region.RegionName == "" {
		return errors.New("region name cannot be empty")
	}
	if len(region.RegionName) > 25 {
		return errors.New("region name cannot exceed 25 characters")
	}
	return s.regionRepo.Update(ctx, region)
}

func (s *regionService) DeleteRegion(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("region ID cannot be empty")
	}
	return s.regionRepo.Delete(ctx, id)
}
