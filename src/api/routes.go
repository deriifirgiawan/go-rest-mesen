package api

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/controllers"
	auth "rest-app-pos/src/modules/auth/routes"
	product "rest-app-pos/src/modules/product/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *containers.AppContainer) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/ping", controllers.Ping)

	auth.AuthRoutes(api, app)
	product.ProductRoutesOwner(api, app)

	return r
}