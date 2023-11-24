package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
)

type (
	BookRepository interface {
		Find(whereClause []query.WhereClause) ([]models.Book, error)
		Save(book *models.Book) (*models.Book, error)
		FindBookByTitle(title string) (*models.Book, error)
	}

	BookUseCase interface {
		GetAllBooks(whereClause []query.WhereClause) ([]models.Book, error)
		CreateBook(book *models.Book) (*models.Book, error)
	}

	BookHandler interface {
		GetAllBooks(c *gin.Context)
		CreateBook(c *gin.Context)
	}
)
