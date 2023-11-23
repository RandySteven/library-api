package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"gorm.io/gorm"
)

type bookUseCase struct {
	repo interfaces.BookRepository
}

// CreateBook implements interfaces.BookUseCase.
func (usecase *bookUseCase) CreateBook(book *models.Book) (*models.Book, error) {
	book, err := usecase.repo.Save(book)
	if err != nil {
		switch err {
		case gorm.ErrDuplicatedKey:
			return nil, gorm.ErrDuplicatedKey
		default:
			return nil, err
		}
	}
	return book, nil
}

// GetAllBooks implements interfaces.BookUseCase.
func (usecase *bookUseCase) GetAllBooks(searchBook *models.Book) ([]models.Book, error) {
	return usecase.repo.Find(searchBook)
}

func NewBookUseCase(repo interfaces.BookRepository) *bookUseCase {
	return &bookUseCase{repo}
}

var _ interfaces.BookUseCase = &bookUseCase{}
