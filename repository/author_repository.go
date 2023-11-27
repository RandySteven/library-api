package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

// FindAuthorByName implements interfaces.AuthorRepository.
func (repo *authorRepository) FindAuthorByName(ctx context.Context, name string) (*models.Author, error) {
	var author *models.Author
	err := repo.db.
		Table("authors").
		Where("author = ?", name).
		Scan(&author).Error
	if err != nil {
		return nil, err
	}
	return author, nil
}

// Find implements interfaces.AuthorRepository.
func (repo *authorRepository) Find(ctx context.Context) ([]models.Author, error) {
	var authors []models.Author
	err := repo.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

// Save implements interfaces.AuthorRepository.
func (repo *authorRepository) Save(ctx context.Context, author *models.Author) (*models.Author, error) {
	err := repo.db.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return author, nil
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{db}
}

var _ interfaces.AuthorRepository = &authorRepository{}
