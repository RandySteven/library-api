package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"not null"`
	Email       string         `gorm:"not null;unique"`
	PhoneNumber string         `gorm:"not null;unique"`
	Password    string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null;default:current_timestamp" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Books       []Book         `gorm:"many2many:borrows;refer:book_id"`
}
