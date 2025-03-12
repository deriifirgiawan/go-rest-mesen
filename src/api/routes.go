package api

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/controllers"
	"rest-app-pos/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *containers.AppContainer) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	{
		api.GET("/ping", controllers.Ping)

		// Auth
		api.POST("/auth/register", app.AuthController.Register)
		api.POST("/auth/login", app.AuthController.Login)

		// Product
		api.POST("/product", middlewares.AuthMiddleware(), app.ProductController.CreateProduct)
		api.PUT("/product", middlewares.AuthMiddleware(), app.ProductController.UpdateProduct)
		api.DELETE("/product/:id", middlewares.AuthMiddleware(), app.ProductController.DeleteProduct)
	}

	return r
}