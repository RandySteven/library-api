package handler

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BorrowHandler struct {
	usecase interfaces.BorrowUseCase
}

// ReturnBorrowBook implements interfaces.BorrowHandler.
func (handler *BorrowHandler) ReturnBorrowBook(c *gin.Context) {
	id := c.Param("id")
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	borrow, err := handler.usecase.ReturnBorrowedBookByBorrowId(ctx, uint(idInt))
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := response.Response{
		Message: "Returned book success",
		Data:    borrow,
	}
	c.JSON(http.StatusOK, resp)
}

// CreateBorrowRecord implements interfaces.BorrowHandler.
func (handler *BorrowHandler) CreateBorrowRecord(c *gin.Context) {
	var request request.BorrowRequest
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
		return
	}

	borrow := &models.Borrow{
		UserID: request.UserID,
		BookID: request.BookID,
	}

	borrowRecord, err := handler.usecase.CreateBorrowRecord(ctx, borrow)
	if err != nil {
		c.Error(err)
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
	var search request.SearchBorrow
	c.ShouldBindQuery(&search)
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	val := reflect.ValueOf(&search).Elem()
	var whereClauses []query.WhereClause
	for i := 0; i < val.NumField(); i++ {
		whereClause := &query.WhereClause{
			Field:     val.Type().Field(i).Name,
			Value:     fmt.Sprintf("%v", val.Field(i).Interface()),
			Condition: enums.Equal,
		}
		whereClauses = append(whereClauses, *whereClause)
	}

	borrows, err := handler.usecase.GetAllBorrowsRecord(ctx, whereClauses)
	if err != nil {
		c.Error(err)
		return
	}
	resp := response.Response{
		Message: "Success get borrow record",
		Data:    borrows,
	}
	c.JSON(http.StatusOK, resp)
}

func NewBorrowHandler(usecase interfaces.BorrowUseCase) *BorrowHandler {
	return &BorrowHandler{usecase: usecase}
}

var _ interfaces.BorrowHandler = &BorrowHandler{}
