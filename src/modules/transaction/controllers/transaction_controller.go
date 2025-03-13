package transaction

import (
	"net/http"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/modules/transaction/services"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
	return &TransactionController{transactionService}
}

func (tc *TransactionController) CreateTransaction(context *gin.Context) {
	var input dto.TransactionRequestDto

	// Validasi input
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := tc.transactionService.CreateTransaction(input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}
