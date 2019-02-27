package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64IsLessThan is a validator object.
type Int64IsLessThan struct {
	Name          string
	Field         int64
	ComparedName  string
	ComparedField int64
}

// Validate adds an error if the Field is not less than the ComparedField.
func (v *Int64IsLessThan) Validate(e *validator.Errors) {
	if v.Field < v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not less than %d", v.Field, v.ComparedField))
}
