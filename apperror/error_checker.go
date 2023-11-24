package apperror

import "net/http"

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
