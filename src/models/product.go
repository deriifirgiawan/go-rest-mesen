package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
	Description string `grom:"type:varchar(255)"`
	Price float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Visible bool `json:"visible" gorm:"default: true"`
	CategoryID uint `json:"category_id"`
	Category *Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MerchantID uint `gorm:"not null"`
	Merchant Merchant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}