package models

type MerchantUser struct {
	MerchantID uint `json:"merchant_id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"primaryKey"`
}