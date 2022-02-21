package models

import "time"

type Product struct {
	ID          uint64    `json:"id" gorm:"primary_key:auto_increment"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Title       string    `json:"title" gorm:"type:varchar(128)"`
	Description string    `json:"description" gorm:"type:text"`
	Price       uint16    `json:"price"`
	Stock       uint16    `json:"stock"`
	StockStatus bool      `json:"stock_status"`
	// CompanyID   uint64    `json:"-" gorm:"not null"`
	// Company     Company   `json:"company" gorm:"foreignkey:CompanyID" constraint:onUpdate:CASCADE, onDelete:CASCADE"`
	UserID uint64 `json:"-" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignkey:UserID" constraint:onUpdate:CASCADE, onDelete:CASCADE"`
}
