package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
	RoleID uint `json:"role_id"`
	Role Role `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Merchant *Merchant `json:"merchant,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
