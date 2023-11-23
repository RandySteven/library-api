package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

// Find implements interfaces.BookRepository.
func (repo *bookRepository) Find() ([]models.Book, error) {
	var books []models.Book
	err := repo.db.Table("books").Scan(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

var _ interfaces.BookRepository = &bookRepository{}
