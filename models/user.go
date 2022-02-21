package models

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key:auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Nickname  string    `json:"nickname" binding:"required" gorm:"type:varchar(32)"`
	Email     string    `json:"email" binding:"required" gorm:"type:varchar(255); uniqueIndex"`
	Password  string    `json:"-"  gorm:"->;<-; not null"`
	Token     string    `json:"token,omitempty" gorm:"-"`
	// CompanyID uint64     `json:"-" gorm:"not null"`
	// Company   Company    `json:"company" gorm:"foreignkey:CompanyID" constraint:onUpdate:CASCADE, onDelete:CASCADE"`
	Products *[]Product `json:"products,omitempty"`
}
