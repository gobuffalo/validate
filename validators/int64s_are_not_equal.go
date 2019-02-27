package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64sAreNotEqual is a validator object.
type Int64sAreNotEqual struct {
	Name          string
	Field         int64
	ComparedName  string
	ComparedField int64
}

// Validate adds an error if the Field is equal to the ComparedField.
func (v *Int64sAreNotEqual) Validate(e *validator.Errors) {
	if v.Field != v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.ComparedField))
}
