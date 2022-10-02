package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host         = "localhost"
	user         = "postgres"
	password     = "postgres"
	databasePort = "5432"
	databaseName = "postgres"
)

func InitDb() *gorm.DB {
	config := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, databaseName, databasePort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database")
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db
}
