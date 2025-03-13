package product

import (
	"rest-app-pos/src/containers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup, app *containers.AppContainer) {
	product := router.Group("/product")

	{
		product.GET("/list", app.ProductController.GetAllProducts)
		product.GET("/categories", app.ProductController.GetAllCategories)
		product.GET("/category/:id", app.ProductController.GetCategoryById)
	}
}