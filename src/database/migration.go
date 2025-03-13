package database

import (
	"log"
	"rest-app-pos/src/models"

	"gorm.io/gorm"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Merchant{}, &models.Category{}, &models.Product{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database successfully migration")

	createProductListView(DB)
}

func createProductListView(db *gorm.DB) {
	log.Println("Create View Product")

	viewSQL := `
	CREATE OR REPLACE VIEW product_list AS
	SELECT 
		p.id AS id,
		p.name AS name,
		p.price,
		p.created_at,
		p.updated_at,
		p.visible,
		c.id AS category_id,
		c.name AS category_name,
		m.id AS merchant_id,
		m.name AS merchant_name
	FROM products p
	LEFT JOIN categories c ON p.category_id = c.id
	LEFT JOIN merchants m ON p.merchant_id = m.id;
	`

	err := db.Exec(viewSQL).Error
	if err != nil {
		log.Fatalf("Failed to create view: %v", err)
	} else {
		log.Println("View 'product_details' created successfully!")
	}
}