package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
)

type bookUseCase struct {
	bookRepo   interfaces.BookRepository
	authorRepo interfaces.AuthorRepository
}

// CreateBook implements interfaces.BookUseCase.
func (usecase *bookUseCase) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	bookExists, _ := usecase.bookRepo.FindBookByTitle(ctx, book.Title)
	if bookExists != nil {
		return nil, apperror.NewErrNoDuplication("books", "title", book.Title).Err
	}
	book, err := usecase.bookRepo.Save(ctx, book)
	return book, err
}

// GetAllBooks implements interfaces.BookUseCase.
func (usecase *bookUseCase) GetAllBooks(ctx context.Context, whereClauses []query.WhereClause) ([]models.Book, error) {
	return usecase.bookRepo.Find(ctx, whereClauses)
}

func NewBookUseCase(
	bookRepo interfaces.BookRepository,
	authorRepo interfaces.AuthorRepository,
) *bookUseCase {
	return &bookUseCase{bookRepo, authorRepo}
}

var _ interfaces.BookUseCase = &bookUseCase{}
