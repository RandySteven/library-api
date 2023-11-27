package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BorrowHandlerTestSuite struct {
	suite.Suite
	borrowUseCase *mocks.BorrowUseCase
	borrowHandler handler.BorrowHandler
	r             *gin.Engine
}

func (suite *BorrowHandlerTestSuite) Setup() {
	suite.borrowUseCase = mocks.NewBorrowUseCase(suite.T())
	suite.borrowHandler = *handler.NewBorrowHandler(suite.borrowUseCase)
	suite.r = gin.Default()
	suite.r.Use(server.ErrorMiddleware())
}

func TestBorrowHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}

func (suite *BorrowHandlerTestSuite) TestCreateBorrowRecord() {
	suite.Run("should return 201 after create borrow record", func() {
		requestBody := `{
			"user_id": 1,
			"book_id": 1
		}`

		suite.borrowUseCase.On("CreateBorrowRecord", mock.AnythingOfType("*models.Borrow")).Return(mock.AnythingOfType("*models.Borrow"), nil)
		req, err := http.NewRequest(http.MethodPost, "/v1/borrowing-record", strings.NewReader(requestBody))
		suite.NoError(err)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.borrowHandler.CreateBorrowRecord(ctx)

		suite.Equal(http.StatusCreated, w.Code)
	})

	suite.Run("should return 500 after while failure to make borrow record", func() {
		suite.borrowUseCase.On("CreateBorrowRecord", mock.AnythingOfType("*models.Borrow")).Return(nil, errors.New("mock error"))
		req, err := http.NewRequest(http.MethodPost, "/v1/borrowing-record", nil)
		suite.NoError(err)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.borrowHandler.CreateBorrowRecord(ctx)

		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}
