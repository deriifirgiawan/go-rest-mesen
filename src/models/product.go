package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255)"`
	Price  float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Quantity uint `json:"quantity" gorm:"not null;default:0"`
	Visible bool `json:"visible" gorm:"default:true"`
	CategoryID uint `json:"category_id,omitempty"`
	Category *Category `json:"category,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MerchantID uint `json:"merchant_id" gorm:"not null"`
	Merchant *Merchant `json:"merchant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
