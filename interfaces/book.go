package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"github.com/gin-gonic/gin"
)

type (
	BookRepository interface {
		Find(searchBook *models.Book) ([]models.Book, error)
		Save(book *models.Book) (*models.Book, error)
	}

	BookUseCase interface {
		GetAllBooks(searchBook *models.Book) ([]models.Book, error)
		CreateBook(book *models.Book) (*models.Book, error)
	}

	BookHandler interface {
		GetAllBooks(c *gin.Context)
		CreateBook(c *gin.Context)
	}
)
