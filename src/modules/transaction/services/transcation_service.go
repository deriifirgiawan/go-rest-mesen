package services

import (
	"errors"
	"rest-app-pos/src/database"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
	"rest-app-pos/src/modules/transaction/repository"
	"rest-app-pos/src/utils"
)

type TransactionService interface {
	CreateTransaction(payload dto.TransactionRequestDto) (*models.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(payload dto.TransactionRequestDto) (*models.Transaction, error) {
	var totalAmount float64
	var transactionDetails []models.TransactionDetail

	for _, product := range payload.Products {
		var p models.Product

		if err := database.DB.First(&p, product.ProductID).Error; err != nil {
			return nil, errors.New("product not found")
		}

		subtotal := p.Price * float64(product.Quantity)
		totalAmount += subtotal

		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
			Subtotal:  subtotal,
		})
	}

	invoiceNumber := "INV-" + utils.GenerateInvoiceNumber()

	transaction := &models.Transaction{
		InvoiceNumber: invoiceNumber,
		UserID:        payload.UserID,
		TotalAmount:   totalAmount,
		PaymentMethod: payload.PaymentMethod,
		Status:        "Pending",
	}

	if err := database.DB.Create(transaction).Error; err != nil {
		return nil, err
	}

	for i := range transactionDetails {
		transactionDetails[i].TransactionID = transaction.ID
	}

	if err := database.DB.Create(&transactionDetails).Error; err != nil {
		return nil, err
	}

	transaction.Details = transactionDetails
	return transaction, nil
}
