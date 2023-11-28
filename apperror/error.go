package apperror

import (
	"errors"
	"fmt"
)

type ErrNoDuplication struct {
	Resource string
	Field    string
	Value    string
	Err      error
}

type ErrBookIdNotFound struct {
	Err error
}

type ErrUserIdNotFound struct {
	Err error
}

type ErrNoColumnExists struct {
	Err error
}

type ErrBookQuantityZero struct {
	Err error
}

type ErrBorrowStatusAlreadyReturned struct {
	Err error
}

type ErrBadRequest struct {
	Err error
}

type ErrUnauthorized struct {
	Err error
}

type ErrBorrowRecordNotFound struct {
	Err error
}

type ErrPermissionDenied struct {
	Err error
}

type ErrPasswordTooLong struct {
	Err error
}

func NewErrNoDuplication(resource, field, value string) *ErrNoDuplication {
	return &ErrNoDuplication{
		Resource: resource,
		Field:    field,
		Value:    value,
		Err:      fmt.Errorf("%s for %s at table %s already exists", value, field, resource),
	}
}

func NewErrBookIdNotFound() *ErrBookIdNotFound {
	return &ErrBookIdNotFound{
		Err: errors.New("book id not found"),
	}
}

func NewErrBadRequest(message string) *ErrBadRequest {
	return &ErrBadRequest{
		Err: fmt.Errorf(message),
	}
}

func NewErrBookQuantityZero() *ErrBookQuantityZero {
	return &ErrBookQuantityZero{
		Err: errors.New("book currently not available"),
	}
}

func NewErrBorrowStatusAlreadyReturned() *ErrBorrowStatusAlreadyReturned {
	return &ErrBorrowStatusAlreadyReturned{
		Err: errors.New("book already returned"),
	}
}

func NewErrUserIdNotFound() *ErrUserIdNotFound {
	return &ErrUserIdNotFound{
		Err: errors.New("user id not found"),
	}
}

func NewErrUnauthorized() *ErrUnauthorized {
	return &ErrUnauthorized{
		Err: errors.New("unauthorized user"),
	}
}

func NewErrBorrowRecordNotFound() *ErrBorrowRecordNotFound {
	return &ErrBorrowRecordNotFound{
		Err: errors.New("borrow record not found"),
	}
}

func NewErrPermissionDenied() *ErrPermissionDenied {
	return &ErrPermissionDenied{
		Err: errors.New("permission denied"),
	}
}

func NewErrPasswordTooLong() *ErrPasswordTooLong {
	return &ErrPasswordTooLong{
		Err: errors.New("Password too long"),
	}
}

func (e *ErrNoDuplication) Error() string {
	return e.Err.Error()
}

func (e *ErrBookIdNotFound) Error() string {
	return e.Err.Error()
}

func (e *ErrBookQuantityZero) Error() string {
	return e.Err.Error()
}

func (e *ErrBorrowStatusAlreadyReturned) Error() string {
	return e.Err.Error()
}

func (e *ErrUserIdNotFound) Error() string {
	return e.Err.Error()
}

func (e *ErrBadRequest) Error() string {
	return e.Err.Error()
}

func (e *ErrUnauthorized) Error() string {
	return e.Err.Error()
}

func (e *ErrBorrowRecordNotFound) Error() string {
	return e.Err.Error()
}

func (e *ErrPermissionDenied) Error() string {
	return e.Err.Error()
}

func (e *ErrPasswordTooLong) Error() string {
	return e.Err.Error()
}

type ErrorWrapper struct {
	ErrType string
	Message string
	Error   error
}

func NewErrorWrapper(errType string, message string, err error) *ErrorWrapper {
	return &ErrorWrapper{}
}
