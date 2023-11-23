package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
)

type bookUseCase struct {
	repo interfaces.BookRepository
}

// GetAllBooks implements interfaces.BookUseCase.
func (usecase *bookUseCase) GetAllBooks() ([]models.Book, error) {
	return usecase.repo.Find()
}

func NewBookUseCase(repo interfaces.BookRepository) *bookUseCase {
	return &bookUseCase{repo}
}

var _ interfaces.BookUseCase = &bookUseCase{}
