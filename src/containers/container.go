package containers

import (
	"rest-app-pos/src/config"
	"rest-app-pos/src/controllers"
	"rest-app-pos/src/database"
	"rest-app-pos/src/repository"
	"rest-app-pos/src/services"
)


type AppContainer struct {
	AuthController *controllers.AuthController
	ProductController *controllers.ProductController
}

func InitAppDependencies() *AppContainer {
	config.LoadConfig()
	database.ConnectDB()

	// Migrate Database
	database.MigrateDB()

	// Seeders
	database.SeedRoles()
	database.SeedCategories()

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)


	productRepo := repository.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	return &AppContainer{
		AuthController: authController,
		ProductController: productController,
	}
}