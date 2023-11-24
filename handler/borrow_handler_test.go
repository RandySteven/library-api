package handler_test

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
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
}
