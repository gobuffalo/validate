package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64sAreNotEqual is a validator object.
type Int64sAreNotEqual struct {
	Name     string
	Field    int64
	Compared int64
}

// Validate adds an error if the field is equal to the compared value.
func (v *Int64sAreNotEqual) Validate(e *validator.Errors) {
	if v.Field != v.Compared {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.Compared))
}
