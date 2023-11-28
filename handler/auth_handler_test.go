package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
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

func (suite *AuthHandlerTestSuite) SetupSubTest() {
	suite.authUseCase = mocks.NewAuthUseCase(suite.T())
	suite.authHandler = handler.NewAuthHandler(suite.authUseCase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

var userAuths = []models.User{
	{
		ID:       1,
		Name:     "Randy Steven",
		Email:    "randy.steven@shopee.com",
		Password: "test_1234",
	},
}

func (suite *AuthHandlerTestSuite) TestLoginAuth() {
	suite.Run("should return 200 when success to login", func() {
		requestBody := `{
			"email": "randysteven12@gmail.com",
			"password": "test_1234"
		}`
		suite.authUseCase.On("LoginUserByEmail", mock.Anything, mock.Anything, mock.Anything).
			Return("token", nil)

		req, _ := http.NewRequest("POST", "/v1/login", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		suite.authHandler.LoginUser(c)
		suite.Equal(http.StatusOK, w.Code)
	})

	suite.Run("should return 400 due error while encrypt password using more than limit password", func() {
		requestBody := `{
			"email": "randysteven12@gmail.com",
			"password": "test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234"
		}`
		errJson := `{
			"errors": "Password too long"
		}`
		var expectedErr interface{}
		json.Unmarshal([]byte(errJson), expectedErr)
		req, _ := http.NewRequest("POST", "/v1/login", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		var actualErr interface{}

		suite.router.POST("/v1/login", suite.authHandler.LoginUser)
		suite.router.ServeHTTP(w, req)
		json.Unmarshal(w.Body.Bytes(), &actualErr)

		suite.Equal(http.StatusBadRequest, w.Code)
		suite.JSONEq(errJson, w.Body.String())
	})

	suite.Run("should return 500 when error in database", func() {
		requestBody := `{
			"email": "randysteven12@gmail.com",
			"password": "test_1234"
		}`
		suite.authUseCase.On("LoginUserByEmail", mock.Anything, mock.Anything, mock.Anything).
			Return("", errors.New("mock error"))

		req, _ := http.NewRequest("POST", "/v1/login", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		suite.router.POST("/v1/login", suite.authHandler.LoginUser)
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})

}

func (suite *AuthHandlerTestSuite) TestRegisterAuth() {
	suite.Run("should return 201 when user succes to register", func() {
		requestBody := `{
			"email": "randysteven12@gmail.com",
			"name": "Randy Steven",
			"phone_number": "0812345678",
			"password": "test_1234"
		}`
		suite.authUseCase.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
			Return(&userAuths[0], nil)

		req, _ := http.NewRequest("POST", "/v1/register", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		suite.authHandler.RegisterUser(c)
		suite.Equal(http.StatusCreated, w.Code)
	})

	suite.Run("should return 400 when user succes to register", func() {
		requestBody := `{
			"email": "randysteven12@gmail.com",
			"name": "Randy Steven",
			"phone_number": "0812345678",
			"password": "test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234test_1234"
		}`

		req, _ := http.NewRequest("POST", "/v1/register", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		suite.router.POST("/v1/register", suite.authHandler.RegisterUser)
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("should return 500 when user succes to register", func() {
		requestBody := `{
			"email": "randysteven12@gmail.com",
			"name": "Randy Steven",
			"password": "test_1234"
		}`
		suite.authUseCase.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
			Return(nil, errors.New("mock error"))

		req, _ := http.NewRequest("POST", "/v1/register", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		suite.router.POST("/v1/register", suite.authHandler.RegisterUser)
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func (suite *AuthHandlerTestSuite) TestLogoutAuth() {
	suite.Run("should return success to logout 200", func() {
		suite.router.POST("/logout", suite.authHandler.LogoutUser)
		req, _ := http.NewRequest(http.MethodPost, "/logout", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
		suite.T().Log("body : ", w.Body)
	})
}
