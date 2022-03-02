package dto

type RegisterDTO struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
