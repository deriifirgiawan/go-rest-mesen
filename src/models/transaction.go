package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	InvoiceNumber string `json:"invoice_number" gorm:"unique;not null"`
	UserID *uint `json:"user_id" gorm:"index"`
	User *User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount float64 `json:"total_amount" gorm:"not null"`
	PaymentMethod string `json:"payment_method" gorm:"type:varchar(100);not null"`
	Status string `jsont:"status" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Details []TransactionDetail `json:"transaction_details" gorm:"foreignKey:TransactionID"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionID uint `json:"transcation_id" gorm:"index;not null"`
	Transaction Transaction `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID uint `json:"product_id" gorm:"index;not null"`
	Product Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity uint `json:"quantity" gorm:"not null"`
	Subtotal float64 `json:"subtotal" gorm:"not null"`
}
