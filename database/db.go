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
	//dsn := "host=localhost user=postgres password=p4ssw0rd dbname=assignment2 port=5432 sslmode=disable"
	dsn := "dbname=d8d9pngqm8aqic host=ec2-34-230-153-41.compute-1.amazonaws.com port=5432 user=bvndtkkudwvpax password=0a025440ac02afeda30594e3e3d06eda3e43183e095442ff6bcdad7695110e81 sslmode=require"

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
