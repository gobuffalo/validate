package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintIsLessThan is a validator object
type UintIsLessThan struct {
	Name          string
	Field         uint
	ComparedName  string
	ComparedField uint
}

// Validate adds an error if the Field is not less than the ComparedField.
func (v *UintIsLessThan) Validate(e *validator.Errors) {
	if v.Field < v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not less than %d", v.Field, v.ComparedField))
}
