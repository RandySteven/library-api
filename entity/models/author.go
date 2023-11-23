package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Books     []Book         `gorm:"foreignKey:author_id;references:id", json:"books,omitempty"`
}
