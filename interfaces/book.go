package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"github.com/gin-gonic/gin"
)

type (
	BookRepository interface {
		Find() ([]models.Book, error)
	}

	BookUseCase interface {
		GetAllBooks() ([]models.Book, error)
	}

	BookHandler interface {
		GetAllBooks(c *gin.Context)
	}
)
