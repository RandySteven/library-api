package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var borrows = []models.Borrow{
	{
		ID:            1,
		UserID:        1,
		BookID:        1,
		BorrowStatus:  "BORROWED",
		BorrowingDate: time.Now(),
	},
	{
		ID:            2,
		UserID:        1,
		BookID:        1,
		BorrowStatus:  "BORROWED",
		BorrowingDate: time.Now(),
	},
}

type BorrowHandlerTestSuite struct {
	suite.Suite
	borrowUseCase *mocks.BorrowUseCase
	borrowHandler *handler.BorrowHandler
	router        *gin.Engine
}

func (suite *BorrowHandlerTestSuite) SetupSubTest() {
	suite.borrowUseCase = mocks.NewBorrowUseCase(suite.T())
	suite.borrowHandler = handler.NewBorrowHandler(suite.borrowUseCase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestBorrowHandler(t *testing.T) {
	suite.Run(t, new(BorrowHandlerTestSuite))
}

func (suite *BorrowHandlerTestSuite) TestCreateBorrowRecord() {
	suite.Run("should return 201 after create borrow record", func() {
		requestBody := `{
			"user_id": 1,
			"book_id": 1
		}`

		suite.borrowUseCase.
			On("CreateBorrowRecord", mock.Anything, mock.AnythingOfType("*models.Borrow")).
			Return(&borrows[0], nil)
		req, err := http.NewRequest(http.MethodPost, "/v1/borrowing-records", strings.NewReader(requestBody))
		suite.NoError(err)
		w := httptest.NewRecorder()
		// c, _ := gin.CreateTestContext(w)

		suite.router.POST("/v1/borrowing-records", suite.borrowHandler.CreateBorrowRecord)
		// suite.borrowHandler.CreateBorrowRecord(c)
		suite.router.ServeHTTP(w, req)
		suite.T().Log("Response : ", w.Body)
		suite.Equal(http.StatusCreated, w.Code)
	})

	suite.Run("should return 500 after while failure to make borrow record", func() {
		requestBody := `{
			"user_id": 1,
			"book_id": 1
		}`

		suite.borrowUseCase.On("CreateBorrowRecord", mock.Anything, mock.AnythingOfType("*models.Borrow")).Return(nil, errors.New("mock error"))

		req, err := http.NewRequest(http.MethodPost, "/v1/borrowing-records", strings.NewReader(requestBody))
		suite.NoError(err)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.router.POST("/v1/borrowing-records", suite.borrowHandler.CreateBorrowRecord)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func (suite *BorrowHandlerTestSuite) TestReturnBook() {
	suite.Run("should return 200 after updated borrow record into return", func() {
		requestBody := `{
			"user_id": 1,
			"book_id": 1
		}`

		suite.borrowUseCase.
			On("ReturnBorrowedBookByBorrowId", mock.Anything, mock.Anything, mock.Anything).
			Return(&borrows[0], nil)
		req, err := http.NewRequest(http.MethodPut, "/v1/borrowing-records/1", strings.NewReader(requestBody))
		suite.NoError(err)
		w := httptest.NewRecorder()
		q := req.URL.Query()
		q.Add("id", "1")

		suite.router.PUT("/v1/borrowing-records/:id", suite.borrowHandler.ReturnBorrowBook)
		suite.router.ServeHTTP(w, req)
		suite.T().Log("Response : ", w.Body)
		suite.Equal(http.StatusOK, w.Code)
	})
}
