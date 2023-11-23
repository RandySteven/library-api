package handler_test

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
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
}

func (suite *BookHandlerTestSuite) SetupSubTest() {
	suite.bookUseCase = mocks.NewBookUseCase(suite.T())
	suite.bookHandler = handler.NewBookHandler(suite.bookUseCase)
	suite.router = gin.Default()
}

func (suite *BookHandlerTestSuite) TestGetAllBooks() {
	suite.Run("should return 200 if get all books success", func() {
		expected, _ := json.Marshal(response.Response{
			Message: "Success get all books",
			Data:    books,
		})
		suite.bookUseCase.On("GetAllBooks", &models.Book{}).Return(books, nil)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		suite.router.GET("/v1/books", suite.bookHandler.GetAllBooks)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	suite.Run("should return 200 if get all books success", func() {
		r := gin.Default()

		suite.bookUseCase.On("GetAllBooks", &models.Book{}).Return(books, nil)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		r.GET("/v1/books", suite.bookHandler.GetAllBooks)
		r.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
	})

	suite.Run("should return 500 while error in query", func() {

		suite.bookUseCase.On("GetAllBooks", &models.Book{}).Return(nil, errors.New("Fake error"))
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		suite.router.GET("/v1/books", suite.bookHandler.GetAllBooks)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusInternalServerError, w.Code)

	})

}

func (suite *BookHandlerTestSuite) TestCreate() {
	suite.Run("should return 201 after inster book", func() {
		// expected, _ := json.Marshal(response.Response{
		// 	Message: "Success get all books",
		// 	Data:    books[0],
		// })

		request := `{
			"name": "Book 1",
			"description": "Book Description 2",
			"quantity": 2,
			"cover": "",
			"author_id": 1
		}`

		suite.bookUseCase.On("CreateBook", mock.AnythingOfType("*models.Book")).Return(books[0], nil)
		req, _ := http.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(request))
		w := httptest.NewRecorder()

		suite.router.POST("/v1/books", suite.bookHandler.CreateBook)
		suite.router.ServeHTTP(w, req)

		log.Println("err : ", w.Body.String())
		suite.Equal(http.StatusCreated, w.Code)
		// suite.Equal(string(expected), w.Body.String())
	})
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}
