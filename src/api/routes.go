package api

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/controllers"
	auth "rest-app-pos/src/modules/auth/routes"
	merchant "rest-app-pos/src/modules/merchant/routes"
	product "rest-app-pos/src/modules/product/routes"
	transaction "rest-app-pos/src/modules/transaction/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *containers.AppContainer) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Allow requests from any domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	api.GET("/ping", controllers.Ping)

	auth.AuthRoutes(api, app)
	product.ProductRoutesOwner(api, app)
	product.ProductRoutes(api, app)
	merchant.MerchantRoutes(api, app)
	transaction.TransactionRoutes(api, app)

	return r
}