package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntIsLessThan is a validator object.
type IntIsLessThan struct {
	Name          string
	Field         int
	ComparedName  string
	ComparedField int
}

// Validate adds an error if the Field is not less than the ComparedField.
func (v *IntIsLessThan) Validate(e *validator.Errors) {
	if v.Field < v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not less than %d", v.Field, v.ComparedField))
}
