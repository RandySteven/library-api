package handler_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	authUseCase *mocks.AuthUseCase
	authHandler *handler.AuthHandler
	router      *gin.Engine
}

func (suite *AuthHandlerTestSuite) Setup() {
	suite.authUseCase = mocks.NewAuthUseCase(suite.T())
	suite.authHandler = handler.NewAuthHandler(suite.authUseCase)
	suite.router = gin.Default()
}

func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (suite *AuthHandlerTestSuite) TestLoginAuth() {
	suite.Run("should return 200 when success to login", func() {

	})
}

func (suite *AuthHandlerTestSuite) TestRegisterAuth() {
	suite.Run("should return 201 when user succes to register", func() {})
}
