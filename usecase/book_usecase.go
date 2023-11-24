package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
)

type bookUseCase struct {
	repo interfaces.BookRepository
}

// CreateBook implements interfaces.BookUseCase.
func (usecase *bookUseCase) CreateBook(book *models.Book) (*models.Book, error) {
	existBook, _ := usecase.repo.FindBookByTitle(book.Title)
	if existBook != nil {
		return nil, apperror.NewErrNoDuplication("books", "title", book.Title).Err
	}
	book, err := usecase.repo.Save(book)
	return book, err
}

// GetAllBooks implements interfaces.BookUseCase.
func (usecase *bookUseCase) GetAllBooks(whereClauses []query.WhereClause) ([]models.Book, error) {
	return usecase.repo.Find(whereClauses)
}

func NewBookUseCase(repo interfaces.BookRepository) *bookUseCase {
	return &bookUseCase{repo}
}

var _ interfaces.BookUseCase = &bookUseCase{}
