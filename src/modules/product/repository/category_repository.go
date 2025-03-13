package repository

import (
	"rest-app-pos/src/database"
	"rest-app-pos/src/models"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryById(id uint) (*models.Category, error)
}

type categoryRepository struct {}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	result := database.DB.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}


func (r *categoryRepository) GetCategoryById(id uint) (*models.Category, error){
	var category models.Category

	if err := database.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}