package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestGetAllBooks(t *testing.T) {
	t.Run("should return 200 if get all books success", func(t *testing.T) {
		bookUseCase := mocks.NewBookUseCase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		r := gin.Default()

		bookUseCase.On("GetAllBooks", &models.Book{}).Return(books, nil)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		r.GET("/v1/books", bookHandler.GetAllBooks)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("should return 200 if get all books success", func(t *testing.T) {
		bookUseCase := mocks.NewBookUseCase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		r := gin.Default()

		bookUseCase.On("GetAllBooks", &models.Book{Title: "Book 1"}).Return(books, nil)
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		r.GET("/v1/books", bookHandler.GetAllBooks)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		bookUseCase := mocks.NewBookUseCase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		r := gin.Default()

		bookUseCase.On("GetAllBooks", &models.Book{}).Return(nil, errors.New("Fake error"))
		req, _ := http.NewRequest(http.MethodGet, "/v1/books", nil)
		w := httptest.NewRecorder()

		r.GET("/v1/books", bookHandler.GetAllBooks)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})

}
