package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Quantity    uint      `gorm:"not null"`
	Cover       string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt   gorm.DeletedAt
}
