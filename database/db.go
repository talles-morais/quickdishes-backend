package database

import (
	"log"

	"github.com/talles-morais/quick-dishes/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("error connecting to database")
	}

	if err := DB.AutoMigrate(&models.Restaurant{}); err != nil {
		log.Panic("migration error, restaurant model")
	}
	if err := DB.AutoMigrate(&models.Client{}); err != nil {
		log.Panic("migration error, client model")
	}
	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Panic("migration error, product model")
	}
	if err := DB.AutoMigrate(&models.Order{}); err != nil {
		log.Panic("migration error, order model")
	}

}
