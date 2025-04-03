package dto

import "rest-app-pos/src/models"

type MerchantUserRequestDto struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	MerchantID uint `json:"merchant_id" binding:"required"`
	UserID uint `json:"user_id"`
}

type ResponseGetListEmploye struct {
	List *[]models.MerchantUserList
	Limit *int
	Page *int
	TotalItems *int64
	TotalPages *int
	Error string
}