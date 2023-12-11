package handler_grpc_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	handler_grpc "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler/grpc"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/status"
)

func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("should return users", func(t *testing.T) {
		usecase := mocks.NewUserUseCase(t)
		handler := handler_grpc.NewUserHandler(usecase)
		user1 := &models.User{
			ID:          1,
			Name:        "Matthew Alfredo",
			Email:       "matthew.alfredo@gmail.com",
			PhoneNumber: "08123456789",
		}

		user2 := &models.User{
			ID:          2,
			Name:        "Randy Steven",
			Email:       "randy.steven@gmail.com",
			PhoneNumber: "08123456789",
		}

		users := []models.User{
			*user1,
			*user2,
		}

		usecase.On("GetAllUsers", mock.Anything, mock.Anything).Return(users, nil)

		res, err := handler.GetAllUsers(context.Background(), nil)

		codes := status.Code(err)

		assert.Equal(t, len(users), len(res.Users))
		assert.Equal(t, code.Code_OK.String(), codes.String())
	})
}
