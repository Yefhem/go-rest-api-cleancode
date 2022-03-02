package dto

type ProductUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint16 `json:"price" form:"price" binding:"required"`
	Stock       uint16 `json:"stock" form:"stock" binding:"required"`
	StockStatus bool   `json:"stock_status" form:"stock_status" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type ProductCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint16 `json:"price" form:"price" binding:"required"`
	Stock       uint16 `json:"stock" form:"stock" binding:"required"`
	StockStatus bool   `json:"stock_status" form:"stock_status" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

/*
{
	"title":"Gezer Erkek Terlik",
	"description":"35-46 numara arası konforlu, plastik, dikişli erkek terliği.",
	"price":50,
	"stock":10,
	"stock_status":true
}
*/
