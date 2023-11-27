package apperror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusBadRequest          = http.StatusBadRequest
	StatusNotFound            = http.StatusNotFound
	StatusInternalServerError = http.StatusInternalServerError
)

type CustomErrorMap struct {
	ErrorMap map[error]int
}

func (custom *CustomErrorMap) InitErrorMap() {
	initErrors(custom)
}

func NewCustomErrorMap() *CustomErrorMap {
	return &CustomErrorMap{
		ErrorMap: make(map[error]int),
	}
}

func initErrors(custom *CustomErrorMap) {
	// 400 - Bad Request - Errors
	custom.ErrorMap[errBadRequest] = StatusBadRequest
	custom.ErrorMap[errNoDuplication] = StatusBadRequest

	// 404 - Not Found - Errors
	custom.ErrorMap[errBookIdNotFound] = StatusNotFound
	custom.ErrorMap[errUserIdNotFound] = StatusNotFound
}

var errBadRequest *ErrBadRequest
var errNoDuplication *ErrNoDuplication
var errBookIdNotFound *ErrBookIdNotFound
var errUserIdNotFound *ErrUserIdNotFound
var errBookQuantityZero *ErrBookQuantityZero
var errBorrowStatusAlreadyReturned *ErrBorrowStatusAlreadyReturned
var errUnauthorized *ErrUnauthorized

func ErrorChecker(c *gin.Context, err error) {
	switch {
	case errors.As(err, &errBadRequest):
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	case errors.As(err, &errNoDuplication):
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	case errors.As(err, &errBookIdNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
	case errors.As(err, &errUserIdNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
	case errors.As(err, &errBookQuantityZero):
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"errors": err.Error()})
	case errors.As(err, &errBorrowStatusAlreadyReturned):
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	case errors.As(err, &errUnauthorized):
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
	}
}
