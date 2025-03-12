package product

import (
	"rest-app-pos/src/containers"
	"rest-app-pos/src/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutesOwner(router *gin.RouterGroup, app *containers.AppContainer) {
	product := router.Group("/owner/product")

	product.Use(middlewares.RoleProtectMiddleware(2))

	{
		product.POST("/", app.ProductController.CreateProduct)
		product.PUT("/", app.ProductController.UpdateProduct)
		product.DELETE("/:id", app.ProductController.DeleteProduct)
	}
}