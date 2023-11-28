package handler_test

import (
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

var users = []models.User{
	{
		ID:          1,
		Name:        "Randy Steven",
		Email:       "randy.steven@shopee.com",
		PhoneNumber: "+6285347391672",
	},
	{
		ID:          2,
		Name:        "Mat Mat",
		Email:       "mat.mat@shopee.com",
		PhoneNumber: "+628123456789",
	},
}

type UserHandlerTestSuite struct {
	suite.Suite
	userUseCase *mocks.UserUseCase
	userHandler *handler.UserHandler
	router      *gin.Engine
}

func (suite *UserHandlerTestSuite) SetupSubTest() {
	suite.userUseCase = mocks.NewUserUseCase(suite.T())
	suite.userHandler = handler.NewUserHandler(suite.userUseCase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestCreateUser() {
	suite.Run("should return 201 if user success created", func() {
		requestBody := `{
			"name": "Randy Steven",
			"email": "randy.steven@shopee.com",
			"phone_number": "+6285347391672"
		}`
		suite.userUseCase.On("CreateUser", mock.Anything, mock.AnythingOfType("*models.User")).Return(&users[0], nil)
		req, err := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(requestBody))
		suite.NoError(err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.userHandler.CreateUser(ctx)

		suite.Equal(http.StatusOK, w.Code)
	})

	suite.Run("should return 400 if user success created", func() {
		requestBody := `{
			"name": "Randy Steven",
		}`
		req, err := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(requestBody))
		suite.NoError(err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.userHandler.CreateUser(ctx)

		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("should return 500 if user success created", func() {
		requestBody := `{
			"name": "Randy Steven",
			"email": "randy.steven@shopee.com",
			"phone_number": "+6285347391672"
		}`
		suite.userUseCase.On("CreateUser", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil, errors.New("mock error"))
		req, err := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(requestBody))
		suite.NoError(err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.userHandler.CreateUser(ctx)

		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func (suite *UserHandlerTestSuite) TestGetAllUsers() {
	suite.Run("should return 200 to get all users", func() {
		suite.userUseCase.On("GetAllUsers", mock.AnythingOfType("[]query.WhereClause")).Return(users, nil)
		req, err := http.NewRequest(http.MethodGet, "/v1/users", nil)
		suite.NoError(err)
		w := httptest.NewRecorder()

		suite.router.GET("/v1/users", suite.userHandler.GetAllUsers)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
	})

	suite.Run("should return 404 even if there is invalid query", func() {
		r := gin.Default()

		suite.router.GET("/v1/users", suite.userHandler.GetAllUsers)
		req, _ := http.NewRequest(http.MethodGet, "/v1/users?invalid=ERR", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		suite.Equal(http.StatusNotFound, w.Code)

	})

	suite.Run("should return 500 due error while process query", func() {
		suite.userUseCase.On("GetAllUsers",
			mock.AnythingOfType("[]query.WhereClause")).
			Return(nil, errors.New("mock error"))
		req, _ := http.NewRequest(http.MethodGet, "/v1/users", nil)
		w := httptest.NewRecorder()

		suite.router.GET("/v1/users", suite.userHandler.GetAllUsers)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}
