package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

/*
{
	"email": "samet@gmail.com",
	"password":"12345"
}
*/
