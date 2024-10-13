package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB()  {
	dsn := "host=localhost user=root password=root dbname=dcs-postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); 
	if err != nil {
		log.Panic("error connecting to database")
	}
	db.AutoMigrate()
}