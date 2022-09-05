package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}