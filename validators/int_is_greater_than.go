package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntIsGreaterThan is a validator object.
type IntIsGreaterThan struct {
	Name          string
	Field         int
	ComparedName  string
	ComparedField int
}

// Validate adds an error if the Field is not greater than the ComparedField.
func (v *IntIsGreaterThan) Validate(e *validator.Errors) {
	if v.Field > v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not greater than %d", v.Field, v.ComparedField))
}
