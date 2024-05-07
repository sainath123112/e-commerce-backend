package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := "host=localhost port= 5432 user=sainath password=Sainath1231 dbname=cart_service sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard.LogMode(logger.Silent),
	})
	return db, err
}
