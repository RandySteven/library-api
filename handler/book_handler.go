package handler

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase interfaces.BookUseCase
}

// CreateBook implements interfaces.BookHandler.
func (handler *BookHandler) CreateBook(c *gin.Context) {
	var request request.BookRequest
	if err := c.ShouldBind(&request); err != nil {
		errorMsg := err.Error()
		errors := strings.Split(errorMsg, "\n")
		resp := response.Response{
			Errors: errors,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	request.Title = strings.Title(request.Title)
	book := &models.Book{
		Title:       request.Title,
		Description: request.Description,
		Quantity:    *request.Quantity,
		Cover:       request.Cover,
		AuthorID:    request.AuthorID,
	}
	book, err := handler.usecase.CreateBook(book)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusConflict, resp)
		return
	}
	bookResponse := response.NewBookResponse(book)
	bookResponse.CreatedAt = &book.CreatedAt
	bookResponse.UpdatedAt = &book.UpdatedAt
	bookResponse.DeletedAt = &book.DeletedAt
	bookResponse.Author = nil
	resp := response.Response{
		Message: "Success get all books",
		Data:    bookResponse,
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAllBooks implements interfaces.BookHandler.
func (handler *BookHandler) GetAllBooks(c *gin.Context) {
	var search request.SearchBook
	if err := c.ShouldBindQuery(&search); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "URL not found")
		return
	}
	book := search.SearchBookToBook()
	books, err := handler.usecase.GetAllBooks(book)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	var bookResponses []response.BookResponse
	for _, book := range books {
		bookResponse := response.NewBookResponse(&book)
		bookResponses = append(bookResponses, *bookResponse)
	}
	resp := response.Response{
		Message: "Success get all books",
		Data:    bookResponses,
	}
	c.JSON(http.StatusOK, resp)
}

func NewBookHandler(usecase interfaces.BookUseCase) *BookHandler {
	return &BookHandler{usecase}
}

var _ interfaces.BookHandler = &BookHandler{}
