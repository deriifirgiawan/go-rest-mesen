package repository

import (
	"log"
	"rest-app-pos/src/database"
	"rest-app-pos/src/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
}

type transactionRepository struct {}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	log.Println(transaction)
	tx := database.DB.Begin()

	// Save Transcation
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, detail := range transaction.Details {
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Model(&models.Product{}).
			Where("id = ?", detail.ProductID).
			Update("quantity", gorm.Expr("quantity - ?", detail.Quantity)).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
