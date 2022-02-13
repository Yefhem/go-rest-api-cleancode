package models

import "time"

// Song represents the songs table in the database
type Song struct {
	ID        uint64     `json:"id" gorm:"primary_key:auto_increment"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:CURRENT_TIMESTAMP"`
	Title     string     `json:"title" gorm:"type:varchar(255)"`
	Artist    string     `json:"artist" gorm:"type:varchar(255)"`
	Album     string     `json:"album" gorm:"type:varchar(255)"`
	UserID    uint64     `json:"-" gorm:"not null"`
	User      User       `json:"user" gorm:"foreignkey:UserID; constraint:onUpdate:CASCADE, onDelete:CASCADE"`
}
