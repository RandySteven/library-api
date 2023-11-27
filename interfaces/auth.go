package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"github.com/gin-gonic/gin"
)

type (
	AuthHandler interface {
		RegisterUser(c *gin.Context)
		LoginUser(c *gin.Context)
		LogoutUser(c *gin.Context)
	}

	AuthRepository interface {
		GetUserByEmail(email string) (*models.User, error)
		RegisterUser(user *models.User) (*models.User, error)
	}

	AuthUseCase interface {
		LoginUserByEmail(email string, password string) (string, error)
		RegisterUser(user *models.User) (*models.User, error)
	}
)
