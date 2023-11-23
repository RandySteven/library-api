package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase interfaces.BookUseCase
}

// GetAllBooks implements interfaces.BookHandler.
func (handler *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := handler.usecase.GetAllBooks()
	if err != nil {
		resp := payloads.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := payloads.Response{
		Message: "Success get all books",
		Data:    books,
	}
	c.JSON(http.StatusOK, resp)
}

func NewBookHandler(usecase interfaces.BookUseCase) *BookHandler {
	return &BookHandler{usecase}
}

var _ interfaces.BookHandler = &BookHandler{}
