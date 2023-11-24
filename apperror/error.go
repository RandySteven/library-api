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

type ErrNoColumnExists struct {
	Err error
}

type ErrBookQuantityZero struct {
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

func NewErrBookQuantityZero() *ErrBookQuantityZero {
	return &ErrBookQuantityZero{
		Err: errors.New("book currently not available"),
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
