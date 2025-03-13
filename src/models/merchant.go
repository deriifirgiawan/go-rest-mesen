package models

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User *User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Products []*Product `json:"products,omitempty" gorm:"foreignKey:MerchantID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
