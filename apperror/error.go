package apperror

import "fmt"

type ErrNoDuplication struct {
	Resource string
	Field    string
	Value    string
}

func (e ErrNoDuplication) Error() string {
	return fmt.Sprintf("%s for %s at table %s already exists", e.Value, e.Field, e.Resource)
}
