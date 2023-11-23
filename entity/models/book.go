package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"not null" json:"description"`
	Quantity    uint           `gorm:"not null" json:"quantity"`
	Cover       string         `json:"cover"`
	CreatedAt   time.Time      `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null;default:current_timestamp" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	AuthorID    uint           `gorm:"not null" json:"author_id"`
	Author      Author         `gorm:"foreignKey:AuthorID;references:ID" json:"author,omitempty"`
}
