package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	BookID        uint           `gorm:"not null" json:"book_id"`
	BorrowStatus  string         `gorm:"not null" json:"borrow_status"`
	BorrowingDate time.Time      `gorm:"not null" json:"borrowing_date"`
	ReturningDate *time.Time     `json:"returning_date"`
	CreatedAt     time.Time      `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"not null;default:current_timestamp" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	User          User           `gorm:"foreignKey:UserID;references:ID"`
	Book          Book           `gorm:"foreignKey:BookID;references:ID"`
}
