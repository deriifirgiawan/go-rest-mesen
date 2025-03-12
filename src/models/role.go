package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	User []User
}