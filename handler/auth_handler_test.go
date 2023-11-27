package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
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
		suite.authUseCase.On("LoginUserByEmail", mock.Anything, mock.Anything, mock.Anything).
			Return("token", nil)

		req, _ := http.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		suite.router.POST("/login", suite.authHandler.LoginUser)
		suite.Equal(http.StatusOK, w.Code)
	})
}

func (suite *AuthHandlerTestSuite) TestRegisterAuth() {
	suite.Run("should return 201 when user succes to register", func() {})
}
