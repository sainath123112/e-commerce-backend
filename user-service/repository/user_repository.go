package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbConnection() (*gorm.DB, error) {
	//dsn := "host=e-commerce-postgres user=sainath password=Sainath1231 dbname=user_service port=5432 sslmode=disable"
	dsn := "host=localhost user=sainath password=Sainath1231 dbname=user_service port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return db, err
}
