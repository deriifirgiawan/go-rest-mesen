package routes

import (
	"rest-app-pos/src/containers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.RouterGroup, app *containers.AppContainer) {
	transactionGroup := router.Group("/transactions")
	{
		transactionGroup.POST("/", app.TransactionController.CreateTransaction)
	}
}
