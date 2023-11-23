package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
)

type bookUseCase struct {
	repo interfaces.BookRepository
}

// CreateBook implements interfaces.BookUseCase.
func (usecase *bookUseCase) CreateBook(book *models.Book) (*models.Book, error) {
	return usecase.repo.Save(book)
}

// GetAllBooks implements interfaces.BookUseCase.
func (usecase *bookUseCase) GetAllBooks(title string) ([]models.Book, error) {
	return usecase.repo.Find(title)
}

func NewBookUseCase(repo interfaces.BookRepository) *bookUseCase {
	return &bookUseCase{repo}
}

var _ interfaces.BookUseCase = &bookUseCase{}
