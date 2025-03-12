package repository

import (
	"rest-app-pos/src/database"
	"rest-app-pos/src/models"
)

type ProductRepository interface {
	FindProductById(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

// Get Product By Id
func (r *productRepository) FindProductById(id uint) (*models.Product, error) {
	var product models.Product

	if err := database.DB.Preload("Category").Preload("User").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}


// Insert Product
func (r *productRepository) Create(product *models.Product) error {
	return database.DB.Create(product).Error
}

// Update Product
func (r *productRepository) Update(product *models.Product) error {
	return database.DB.Save(product).Error
}

// Delete Product By Id
func (r *productRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}