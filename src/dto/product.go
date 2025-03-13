package dto

import "rest-app-pos/src/models"

type ProductRequestDto struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price string `json:"price" binding:"required"`
	CategoryID int `json:"category_id" binding:"required"`
	Quantity uint `json:"quantity" binding:"required"`
}

type ProductRequestUpdateDto struct {
	ID uint `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price string `json:"price" binding:"required"`
	CategoryID int `json:"category_id" binding:"required"`
	Quantity uint `json:"quantity" binding:"required"`
}

type ProductQueryFilterDto struct {
	MerchantID *uint `form:"merchant_id"`
	CategoryID *uint `form:"category_id"`
	Search *string `form:"search"`
	Page *int `form:"page"`
	Limit *int `form:"limit"`
}

type ResponseGetListProduct struct {
	List *[]models.ProductList
	Limit *int
	Page *int
	TotalItems *int64
	TotalPages *int
	Error string
}