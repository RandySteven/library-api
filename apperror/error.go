package apperror

import "fmt"

type ErrNoDuplication struct {
	Resource string
	Field    string
	Value    string
	Err      error
}

func NewErrNoDuplication(resource, field, value string) *ErrNoDuplication {
	return &ErrNoDuplication{
		Resource: resource,
		Field:    field,
		Value:    value,
		Err:      fmt.Errorf("%s for %s at table %s already exists", value, field, resource),
	}
}

func (e *ErrNoDuplication) Error() string {
	return e.Err.Error()
}
