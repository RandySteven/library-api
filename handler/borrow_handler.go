package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"github.com/gin-gonic/gin"
)

type BorrowHandler struct {
	usecase interfaces.BorrowUseCase
}

// CreateBorrowRecord implements interfaces.BorrowHandler.
func (handler *BorrowHandler) CreateBorrowRecord(c *gin.Context) {
	var request request.BorrowRequest
	if err := c.ShouldBind(&request); err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	borrow := &models.Borrow{
		UserID: request.UserID,
		BookID: request.BookID,
	}

	borrowRecord, err := handler.usecase.CreateBorrowRecord(borrow)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.Response{
		Message: "Success add borrow record",
		Data:    borrowRecord,
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllBorrowsRecord implements interfaces.BorrowHandler.
func (handler *BorrowHandler) GetAllBorrowsRecord(c *gin.Context) {
}

func NewBorrowHandler(usecase interfaces.BorrowUseCase) *BorrowHandler {
	return &BorrowHandler{usecase: usecase}
}

var _ interfaces.BorrowHandler = &BorrowHandler{}
