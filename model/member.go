package model

type Member struct {
	Name  string `json:"name" binding:"required" validate:"is_keemmer"`
	Phone int    `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
