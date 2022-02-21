package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
