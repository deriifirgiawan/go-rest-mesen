package database

import (
	"log"
	"rest-app-pos/src/models"

	"gorm.io/gorm"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Merchant{}, &models.Category{}, &models.Product{}, &models.Transaction{}, &models.TransactionDetail{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database successfully migration")

	createProductListView(DB)
	createTransactionListView(DB)
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

func createTransactionListView(db *gorm.DB) {
	log.Println("Create View Transaction")

	viewSQL := `
		CREATE OR REPLACE VIEW transaction_report AS
		SELECT t.id AS transaction_id, t.invoice_number, t.total_amount, 
		t.payment_method, t.status, u.name AS customer_name, 
		p.name AS product_name, td.quantity, td.subtotal
		FROM transactions t
		LEFT JOIN users u ON t.user_id = u.id
		JOIN transaction_details td ON t.id = td.transaction_id
		JOIN products p ON td.product_id = p.id;
	`

	err := db.Exec(viewSQL).Error
	if err != nil {
		log.Fatalf("Failed to create view: %v", err)
	} else {
		log.Println("View 'transaction_report' created successfully!")
	}
}