package testing

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/repositories"
	"github.com/codeid/hr-api-simple/internal/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 1. initial mockup db
func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

// 2. test service create region
func TestRegionService_CreateRegion_Success(t *testing.T) {
	//1.call mockupdb
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	//2. init repos & service
	regionRepo := repositories.NewRegionRepository(gormDB)
	regionService := services.NewRegionService(regionRepo)
	ctx := context.Background()

	//3.assume data come from postman
	region := &models.Region{
		RegionName: "Africa",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`^INSERT INTO "hr"."regions"`).
		WithArgs("Africa").
		WillReturnRows(sqlmock.NewRows([]string{"region_id"}).AddRow(1))
	mock.ExpectCommit()

	err = regionService.CreateRegion(ctx, region)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegionService_UpdateRegion_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	repo := repositories.NewRegionRepository(gormDB)
	service := services.NewRegionService(repo)
	ctx := context.Background()

	region := &models.Region{
		RegionID:   1,
		RegionName: "Asia Updated",
	}

	// Expect transaction dan UPDATE query
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "hr"."regions"`).
		WithArgs("Asia Updated", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = service.UpdateRegion(ctx, region)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegionService_UpdateRegion_ValidationError(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	repo := repositories.NewRegionRepository(gormDB)
	service := services.NewRegionService(repo)
	ctx := context.Background()

	region := &models.Region{
		RegionID:   1,
		RegionName: "", // region name kosong, harus diisi
	}

	// Tidak ada ekspektasi query karena validasi gagal
	err = service.UpdateRegion(ctx, region)
	assert.Error(t, err)
	assert.Equal(t, "region name cannot be empty", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegionService_UpdateRegion_DatabaseError(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	repo := repositories.NewRegionRepository(gormDB)
	service := services.NewRegionService(repo)
	ctx := context.Background()

	region := &models.Region{
		RegionID:   1,
		RegionName: "Asia",
	}

	// Simulasikan database error
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "hr"."regions"`).
		WithArgs("Asia", 1).
		WillReturnError(errors.New("database error"))
	mock.ExpectRollback() // Rollback karena error

	err = service.UpdateRegion(ctx, region)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegionService_DeleteRegion_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	repo := repositories.NewRegionRepository(gormDB)
	service := services.NewRegionService(repo)
	ctx := context.Background()

	regionID := uint(1)

	// Expect transaction dan DELETE query
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "hr"."regions"`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = service.DeleteRegion(ctx, regionID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegionService_DeleteRegion_NotFound(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	repo := repositories.NewRegionRepository(gormDB)
	service := services.NewRegionService(repo)
	ctx := context.Background()

	regionID := uint(999)

	// Simulasikan region tidak ditemukan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "hr"."regions"`).
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	err = service.DeleteRegion(ctx, regionID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegionService_GetRegionByID_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	repo := repositories.NewRegionRepository(gormDB)
	service := services.NewRegionService(repo)
	ctx := context.Background()

	regionID := uint(1)

	// Expect SELECT query
	mock.ExpectQuery(`SELECT * FROM "hr"."regions"`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"region_id", "region_name"}).
			AddRow(1, "Asia"))

	region, err := service.GetRegionByID(ctx, regionID)
	assert.NoError(t, err)
	assert.NotNil(t, region)
	assert.Equal(t, uint(1), region.RegionID)
	assert.Equal(t, "Asia", region.RegionName)
	assert.NoError(t, mock.ExpectationsWereMet())
}
