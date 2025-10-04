package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=hr_db_simple port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	//2. create connection to postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database : %w", err)
	}

	//set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	//set connection pool
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Conected to database successfully")
	return db, nil

}
