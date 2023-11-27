package interfaces

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"github.com/gin-gonic/gin"
)

type (
	AuthorRepository interface {
		Save(ctx context.Context, author *models.Author) (*models.Author, error)
		Find(ctx context.Context) ([]models.Author, error)
		FindAuthorByName(ctx context.Context, name string) (*models.Author, error)
	}
	AuthorUseCase interface {
		GetAllAuthors(ctx context.Context) ([]models.Author, error)
		CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error)
	}
	AuthorHandler interface {
		CreateAuthor(c *gin.Context)
		GetAllAuthors(c *gin.Context)
	}
)
