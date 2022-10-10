package database

import (
	"Assignment2/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "host=localhost user=postgres password=p4ssw0rd dbname=assignment2 port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connection Database : ", err)
	}

	db.Debug().AutoMigrate(models.Item{})
	db.AutoMigrate(&models.Order{})

}

func GetDB() *gorm.DB {
	return db
}
