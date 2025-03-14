package dto

type AuthRequestRegisterDto struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRequestLoginDto struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}