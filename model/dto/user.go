package dto

type UserCreateDTO struct {
	Username string `json:"account" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}
