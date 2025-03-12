package dto

type ProductRequestDto struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price string `json:"price" binding:"required"`
	CategoryID int `json:"category_id" binding:"required"`
}