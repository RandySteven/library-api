package interfaces

import (
	"context"

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
		GetUserByEmail(ctx context.Context, email string) (*models.User, error)
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	}

	AuthUseCase interface {
		LoginUserByEmail(ctx context.Context, email string, password string) (string, error)
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	}
)
