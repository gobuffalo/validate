package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64sAreEqual is a validator object.
type Int64sAreEqual struct {
	Name          string
	Field         int64
	ComparedName  string
	ComparedField int64
}

// Validate add an error if the Field is not equal to the ComparedField.
func (v *Int64sAreEqual) Validate(e *validator.Errors) {
	if v.Field == v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.ComparedField))
}
