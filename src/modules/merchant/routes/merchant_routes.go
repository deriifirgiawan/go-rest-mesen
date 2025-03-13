package routes

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/middlewares"

	"github.com/gin-gonic/gin"
)

func MerchantRoutes(router *gin.RouterGroup, app *containers.AppContainer) {
	product := router.Group("/merchant")

	product.Use(middlewares.RoleProtectMiddleware(2))

	{
		product.GET("/", app.MerchantController.GetMerchant)
		product.POST("/", app.MerchantController.AddMerchant)
		product.PUT("/", app.MerchantController.UpdateMerchant)
	}
}