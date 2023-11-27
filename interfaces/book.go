package interfaces

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
)

type (
	BookRepository interface {
		Find(ctx context.Context, whereClause []query.WhereClause) ([]models.Book, error)
		Save(ctx context.Context, book *models.Book) (*models.Book, error)
		FindBookByTitle(ctx context.Context, title string) (*models.Book, error)
	}

	BookUseCase interface {
		GetAllBooks(ctx context.Context, whereClause []query.WhereClause) ([]models.Book, error)
		CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	}

	BookHandler interface {
		GetAllBooks(c *gin.Context)
		CreateBook(c *gin.Context)
	}
)
