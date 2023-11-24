package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
)

type (
	UserRepository interface {
		Save(user *models.User) (*models.User, error)
		Find(whereClause []query.WhereClause) ([]models.User, error)
	}

	UserUseCase interface {
		CreateUser(user *models.User) (*models.User, error)
		GetAllUsers(whereClause []query.WhereClause) ([]models.User, error)
	}

	UserHandler interface {
		CreateUser(c *gin.Context)
		GetAllUsers(c *gin.Context)
	}
)
