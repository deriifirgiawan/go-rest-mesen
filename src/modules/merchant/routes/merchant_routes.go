package routes

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/middlewares"

	"github.com/gin-gonic/gin"
)

func MerchantRoutes(router *gin.RouterGroup, app *containers.AppContainer) {
	merchant := router.Group("/merchant")

	merchant.Use(middlewares.RoleProtectMiddleware(2))

	{
		merchant.GET("/", app.MerchantController.GetMerchant)
		merchant.POST("/", app.MerchantController.AddMerchant)
		merchant.PUT("/", app.MerchantController.UpdateMerchant)
		merchant.POST("/add-employee", app.MerchantController.AddEmployee)
		merchant.GET("/employees", app.MerchantController.GetAllEmployee)
	}
}