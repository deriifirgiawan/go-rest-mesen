package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID uint `json:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}