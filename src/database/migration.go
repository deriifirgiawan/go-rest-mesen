package database

import (
	"log"
	"rest-app-pos/src/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Merchant{}, &models.Category{}, &models.Product{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database successfully migration")
}