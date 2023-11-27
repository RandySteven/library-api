package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var books = []models.Book{
	{
		ID:          1,
		Title:       "Book 1",
		Description: "Book Description 1",
		Quantity:    2,
		Cover:       "",
	},
	{
		ID:          2,
		Title:       "Book 2",
		Description: "Book Description 2",
		Quantity:    2,
		Cover:       "",
	},
}

type BookHandlerTestSuite struct {
	suite.Suite
	bookUseCase *mocks.BookUseCase
	bookHandler *handler.BookHandler
	router      *gin.Engine
	handler     server.Handlers
}

func (suite *BookHandlerTestSuite) SetupSubTest() {
	suite.bookUseCase = mocks.NewBookUseCase(suite.T())
	suite.bookHandler = handler.NewBookHandler(suite.bookUseCase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func (suite *BookHandlerTestSuite) TestGetAllBooks() {
	suite.Run("should return 200 if get all books success", func() {
		suite.bookUseCase.On("GetAllBooks", mock.AnythingOfType("[]query.WhereClause")).Return(books, nil)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		suite.router.GET("/v1/books", suite.bookHandler.GetAllBooks)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
		suite.NotNil(w.Body)
	})

	suite.Run("should return 200 if get all books success filter by query", func() {
		r := gin.Default()

		suite.bookUseCase.On("GetAllBooks", mock.AnythingOfType("[]query.WhereClause")).Return(books, nil)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		r.GET("/v1/books", suite.bookHandler.GetAllBooks)
		r.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
	})

	suite.Run("should return 400 invalid query", func() {
		r := gin.Default()

		suite.router.GET("/v1/books", suite.bookHandler.GetAllBooks)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books?invalid=ERR", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		suite.Equal(http.StatusNotFound, w.Code)

	})

	suite.Run("should return 500 while error in query", func() {

		suite.bookUseCase.On("GetAllBooks", mock.AnythingOfType("[]query.WhereClause")).Return(nil, errors.New("Fake error"))
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		suite.router.GET("/v1/books", suite.bookHandler.GetAllBooks)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusInternalServerError, w.Code)

	})

}

func (suite *BookHandlerTestSuite) TestCreate() {
	suite.Run("should return 201 after inserting book", func() {
		requestBody := `{
			"title": "Book 1",
			"description": "Book Description 2",
			"quantity": 2,
			"cover": "",
			"author_id": 1
		}`

		suite.bookUseCase.On("CreateBook", mock.AnythingOfType("*models.Book")).Return(&books[0], nil)

		req, err := http.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(requestBody))
		suite.NoError(err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.bookHandler.CreateBook(ctx)

		suite.Equal(http.StatusCreated, w.Code)

		var resp response.Response
		suite.NoError(json.Unmarshal(w.Body.Bytes(), &resp))
	})

	suite.Run("should return 400 bad request while inserting book", func() {
		request := `{
			"title": "Book 1",
			"description": "Book Description 2",
			"quantity": 2,
			"cover": "",
			"author_id": 1
		}`

		// suite.bookUseCase.On("CreateBook", mock.AnythingOfType("*models.Book")).Return(books[0], nil)

		req, err := http.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(request))
		suite.NoError(err)

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.router.POST("/v1/books", suite.bookHandler.CreateBook)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusBadRequest, w.Code)

		var resp response.Response
		suite.NoError(json.Unmarshal(w.Body.Bytes(), &resp))

	})
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}
