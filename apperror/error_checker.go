package apperror

type BadRequest struct {
	Type int
	ErrBadRequest
	ErrNoDuplication
}

type NotFound struct {
	Type int
	ErrBookIdNotFound
	ErrUserIdNotFound
}

var errBadRequest *ErrBadRequest
var errNoDuplication *ErrNoDuplication
var errBookIdNotFound *ErrBookIdNotFound
var errUserIdNotFound *ErrUserIdNotFound

var badRequests []error
var notFounds []error

func initErrors() {
	//400 - Bad Request - Errors
	badRequests = append(badRequests, errBadRequest)
	badRequests = append(badRequests, errNoDuplication)

	//404 - Not Found - Errors
	notFounds = append(notFounds, errBookIdNotFound)
	notFounds = append(notFounds, errUserIdNotFound)
}

var responseMap = make(map[int][]error)

func initErrorMap() {
	responseMap[400] = badRequests
	responseMap[404] = notFounds
}
