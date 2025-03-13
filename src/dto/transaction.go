package dto

type TransactionRequestDto struct {
	UserID        *uint                      `json:"user_id,omitempty"`
	PaymentMethod string                     `json:"payment_method" binding:"required"`
	Products      []TransactionProductDto     `json:"products" binding:"required"`
}

type TransactionProductDto struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}
