package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"github.com/gin-gonic/gin"
)

type (
	AuthorRepository interface {
		Save(author *models.Author) (*models.Author, error)
		Find() ([]models.Author, error)
		FindAuthorByName(name string) (*models.Author, error)
	}
	AuthorUseCase interface {
		GetAllAuthors() ([]models.Author, error)
		CreateAuthor(author *models.Author) (*models.Author, error)
	}
	AuthorHandler interface {
		CreateAuthor(c *gin.Context)
		GetAllAuthors(c *gin.Context)
	}
)
