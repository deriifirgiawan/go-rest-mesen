package database

import (
	"log"
	"rest-app-pos/src/models"
)

func SeedRoles() {
	roles := []models.Role {
		{Name: "ADMIN"},
		{Name: "OWNER"},
		{Name: "EMPLOYEE"},
	}

	for _, role := range roles {
		result := DB.FirstOrCreate(&role, models.Role{Name: role.Name})
		if result.Error != nil {
			log.Printf("Failed to insert role %s: %v", role.Name, result.Error)
		} else if result.RowsAffected > 0 {
			log.Println("Inserted role:", role.Name)
		} else {
			log.Println("Role already exists:", role.Name)
		}
	}
}

func SeedCategories() {
	categories := []models.Category {
		{Name: "Makanan Ringan"},
		{Name: "Makanan Pembuka"},
		{Name: "Minuman Dingin"},
		{Name: "Minuman Panas"},
		{Name: "Minuman Soda"},
		{Name: "Minuman Alkohol"},
		{Name: "Cemilan"},
	}

	for _, category := range categories {
		result := DB.FirstOrCreate(&category, models.Category{Name: category.Name})

		if result.Error != nil {
			log.Printf("Failed to insert role %s: %v", category.Name, result.Error)
		} else if result.RowsAffected > 0 {
			log.Println("Inserted Category:", category.Name)
		} else {
			log.Println("Category already exists:", category.Name)
		}
	}
}