package dto

type MerchantRequestDto struct {
	Name string `json:"name" binding:"required"`
}