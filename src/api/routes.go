package api

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/controllers"
	auth "rest-app-pos/src/modules/auth/routes"
	merchant "rest-app-pos/src/modules/merchant/routes"
	product "rest-app-pos/src/modules/product/routes"
	transaction "rest-app-pos/src/modules/transaction/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *containers.AppContainer) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/ping", controllers.Ping)

	auth.AuthRoutes(api, app)
	product.ProductRoutesOwner(api, app)
	product.ProductRoutes(api, app)
	merchant.MerchantRoutes(api, app)
	transaction.TransactionRoutes(api, app)

	return r
}