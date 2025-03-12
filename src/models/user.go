package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	RoleID uint
	Role Role
	Merchant *Merchant `gorm:"foreignKey:UserID"`
}