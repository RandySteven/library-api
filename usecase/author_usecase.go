package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
)

type authorUseCase struct {
	repo interfaces.AuthorRepository
}

// CreateAuthor implements interfaces.AuthorUseCase.
func (usecase *authorUseCase) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	return usecase.repo.Save(ctx, author)
}

// GetAllAuthors implements interfaces.AuthorUseCase.
func (usecase *authorUseCase) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	return usecase.repo.Find(ctx)
}

func NewAuthorUseCase(repo interfaces.AuthorRepository) *authorUseCase {
	return &authorUseCase{repo: repo}
}

var _ interfaces.AuthorUseCase = &authorUseCase{}
