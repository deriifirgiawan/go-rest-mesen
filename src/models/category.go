package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Products []*Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
