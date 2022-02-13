package models

import "time"

// User represents the user table in the database
type User struct {
	ID        uint64    `json:"id" gorm:"primary_key:auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255); uniqueIndex"`
	Password  string    `json:"-" gorm:"->;<-; not null"`
	Token     string    `json:"token,omitempty" gorm:"-"`
	Songs     *[]Song   `json:"songs,omitempty"`
}
