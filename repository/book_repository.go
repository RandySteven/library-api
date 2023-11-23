package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

// Save implements interfaces.BookRepository.
func (repo *bookRepository) Save(book *models.Book) (*models.Book, error) {
	err := repo.db.Create(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

// Find implements interfaces.BookRepository.
func (repo *bookRepository) Find(searchBook *models.Book) ([]models.Book, error) {
	var books []models.Book
	table := repo.db.Table("books")
	if searchBook != nil {
		table = table.Where(searchBook)
	}
	err := table.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

var _ interfaces.BookRepository = &bookRepository{}
