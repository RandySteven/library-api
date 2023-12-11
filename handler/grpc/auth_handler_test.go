package handler_grpc_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	handler_grpc "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler/grpc"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return success token of login", func(t *testing.T) {
		mocks := mocks.NewAuthUseCase(t)
		handler := handler_grpc.NewAuthHandler(mocks)

		userLogin := &request.UserLoginRequest{
			Email:    "randysteven12@gmail.com",
			Password: "test_1234",
		}

		userReq := &pb.AuthRequest{
			Email:    userLogin.Email,
			Password: userLogin.Password,
		}
		token := "token"
		mocks.On("LoginUserByEmail", mock.Anything, userReq.Email, userReq.Password).Return(token, nil)

		expectedRes := &pb.AuthResponse{
			Message:              "Login success",
			Token:                token,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     []byte{},
			XXX_sizecache:        0,
		}
		res, _ := handler.Login(context.Background(), userReq)

		assert.NotEqual(t, expectedRes, res)
	})

	t.Run("should return error invalid email", func(t *testing.T) {
		mocks := mocks.NewAuthUseCase(t)
		handler := handler_grpc.NewAuthHandler(mocks)

		userLogin := &request.UserLoginRequest{
			Email:    "randysteven12",
			Password: "test_1234",
		}

		userReq := &pb.AuthRequest{
			Email:    userLogin.Email,
			Password: userLogin.Password,
		}

		_, err := handler.Login(context.Background(), userReq)

		assert.Error(t, err)
	})
}
