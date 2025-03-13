package repository

import (
	"errors"
	"fmt"
	"rest-app-pos/src/database"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
)

type ProductRepository interface {
	GetAllProductByMerchant(filter dto.ProductQueryFilterDto) ([]models.ProductList, int64, error)
	FindProductById(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product, merchant_id uint) error
	Delete(id uint) error
}

type productRepository struct {}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

// Get Product By Id
func (r *productRepository) FindProductById(id uint) (*models.Product, error) {
	var product models.Product

	if err := database.DB.Preload("Category").Preload("Merchant").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

// Insert Product
func (r *productRepository) Create(product *models.Product) error {
	return database.DB.Create(product).Error
}

// Update Product
func (r *productRepository) Update(product *models.Product, merchant_id uint) error {
	fmt.Println(product.ID)
	result := database.DB.
		Where("id = ? AND merchant_id = ?", product.ID, merchant_id).
		Updates(product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found or unauthorized to update")
	}

	return nil
}

// Delete Product By Id
func (r *productRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}

func (r *productRepository) GetAllProductByMerchant(filter dto.ProductQueryFilterDto) ([]models.ProductList, int64, error) {
	productList := []models.ProductList{}

	if filter.MerchantID == nil {
		return productList, 0, nil
	}

	db := database.DB.Table("product_list").Where("merchant_id = ?", *filter.MerchantID)

	if filter.CategoryID != nil {
		db = db.Where("category_id = ?", *filter.CategoryID)
	}

	if filter.Search != nil && *filter.Search != "" {
		db = db.Where("LOWER(name) LIKE LOWER(?)", "%"+*filter.Search+"%")
	}

	limit := 10
	page := 1

	if filter.Limit != nil && *filter.Limit > 0 {
		limit = *filter.Limit
	}

	if filter.Page != nil && *filter.Page > 0 {
		page = *filter.Page
	}

	offset := (page - 1) * limit

	var totalItems int64
	db.Count(&totalItems)

	result := db.Limit(limit).Offset(offset).Find(&productList)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return productList, totalItems, nil
}