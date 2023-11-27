package repository

import (
	"context"
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

// FindBookByTitle implements interfaces.BookRepository.
func (repo *bookRepository) FindBookByTitle(ctx context.Context, title string) (*models.Book, error) {
	var book *models.Book
	err := repo.db.Table("books").Where("title = ?", title).Scan(&book).Error
	log.Println(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// Save implements interfaces.BookRepository.
func (repo *bookRepository) Save(ctx context.Context, book *models.Book) (*models.Book, error) {
	err := repo.db.Create(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

// Find implements interfaces.BookRepository.
func (repo *bookRepository) Find(ctx context.Context, whereClauses []query.WhereClause) ([]models.Book, error) {
	var books []models.Book
	table := repo.db.Model(&models.Book{}).
		Preload("Author")
	for _, clause := range whereClauses {
		query := fmt.Sprintf("%s %s ?", clause.Field, clause.Condition)
		if len(clause.Value) > 1 {
			table = table.Where(query, "%"+clause.Value+"%")
		}

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
