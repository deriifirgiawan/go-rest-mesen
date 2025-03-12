package repository

import (
	"rest-app-pos/src/database"
	"rest-app-pos/src/models"
)

type MerchantRepository interface {
	FindById(id uint) (*models.Merchant, error)
	FindByUserId(UserID uint) (*models.Merchant, error)
}

type merchantRepository struct {}

func NewMerchantRepository() MerchantRepository {
	return &merchantRepository{}
}

func (r *merchantRepository) FindById(id uint) (*models.Merchant, error) {
	var merchant models.Merchant

	if err := database.DB.Where("id = ?", id).First(&merchant).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (r *merchantRepository) FindByUserId(UserID uint) (*models.Merchant, error){
	var merchant models.Merchant

	if err := database.DB.Where("user_id = ?", UserID).First(&merchant).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}