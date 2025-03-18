package database

import (
	"fmt"
	"log"
	"todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=admin password=aham dbname=todo port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database", err)
	}
	dbErr := DB.AutoMigrate(&models.User{}, &models.Task{})
	if dbErr != nil {
		log.Fatal("Error while Migrating", err)
	}
	fmt.Println("Database connected")
}
