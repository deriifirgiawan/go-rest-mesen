package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	ID uint `json:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	UserID uint `gorm:"uniqueIndex;not null"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Products []Product `gorm:"foreignKey:MerchantID"`
}