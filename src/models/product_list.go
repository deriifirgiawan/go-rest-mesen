package models

type ProductList struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	CategoryID uint `json:"category_id"`
	CategoryName string `json:"category_name"`
	MerchantID uint `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	Visible bool `json:"visible"`
}