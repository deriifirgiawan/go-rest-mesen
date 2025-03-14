package containers

import (
	"rest-app-pos/src/config"
	"rest-app-pos/src/database"
	auth "rest-app-pos/src/modules/auth/controllers"
	merchant "rest-app-pos/src/modules/merchant/controllers"
	merchantRepo "rest-app-pos/src/modules/merchant/repository"
	merchantService "rest-app-pos/src/modules/merchant/services"
	product "rest-app-pos/src/modules/product/controllers"
	productRepository "rest-app-pos/src/modules/product/repository"
	"rest-app-pos/src/modules/product/services"
	transaction "rest-app-pos/src/modules/transaction/controllers"
	transactionRepo "rest-app-pos/src/modules/transaction/repository"
	transactinService "rest-app-pos/src/modules/transaction/services"
	globalRepository "rest-app-pos/src/repository"
	globalService "rest-app-pos/src/services"
)


type AppContainer struct {
	AuthController *auth.AuthController
	ProductController *product.ProductController
	MerchantController *merchant.MerchantController
	TransactionController *transaction.TransactionController
}

func InitAppDependencies() *AppContainer {
	config.LoadConfig()
	database.ConnectDB()

	// Migrate Database
	database.MigrateDB()

	// Seeders
	database.SeedRoles()
	database.SeedCategories()

	userRepo := globalRepository.NewUserRepository()
	userService := globalService.NewUserService(userRepo)
	authController := auth.NewAuthController(userService)


	productRepo := productRepository.NewProductRepository()
	categoryRepo := productRepository.NewCategoryRepository()
	productService := services.NewProductService(productRepo, categoryRepo)
	productController := product.NewProductController(productService)

	merchantRepo := merchantRepo.NewMerchantRepository()
	merchantService := merchantService.NewMerchantService(merchantRepo)
	merchantController := merchant.NewMerchantController(merchantService)

	transactionRepo := transactionRepo.NewTransactionRepository()
	transactionService := transactinService.NewTransactionService(transactionRepo)
	transactionController := transaction.NewTransactionController(transactionService)

	return &AppContainer{
		AuthController: authController,
		ProductController: productController,
		MerchantController: merchantController,
		TransactionController: transactionController,
	}
}