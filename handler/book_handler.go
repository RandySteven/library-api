package handler

import (
	"log"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
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
	var request *request.BookRequest
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
	}
	book, err := handler.usecase.CreateBook(book)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			errNoDuplication := apperror.ErrNoDuplication{
				Resource: "books",
				Field:    "title",
				Value:    request.Title,
			}
			resp := response.Response{
				Errors: []string{errNoDuplication.Error()},
			}
			c.AbortWithStatusJSON(http.StatusConflict, resp)
			return
		}
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := response.Response{
		Message: "Success get all books",
		Data:    book,
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAllBooks implements interfaces.BookHandler.
func (handler *BookHandler) GetAllBooks(c *gin.Context) {
	var search request.SearchBook
	if err := c.ShouldBindQuery(&search); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad request")
		return
	}
	log.Println("title : ", search.Title)
	book := &models.Book{
		Title:       search.Title,
		Description: search.Description,
		Cover:       search.Cover,
		Quantity:    search.Quantity,
	}
	log.Println(book)
	books, err := handler.usecase.GetAllBooks(book)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := response.Response{
		Message: "Success get all books",
		Data:    books,
	}
	c.JSON(http.StatusOK, resp)
}

func NewBookHandler(usecase interfaces.BookUseCase) *BookHandler {
	return &BookHandler{usecase}
}

var _ interfaces.BookHandler = &BookHandler{}
