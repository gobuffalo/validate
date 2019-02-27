package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintIsGreaterThan is a validator object
type UintIsGreaterThan struct {
	Name          string
	Field         uint
	ComparedName  string
	ComparedField uint
}

// Validate adds an error if the field is not greater than the compared value.
func (v *UintIsGreaterThan) Validate(e *validator.Errors) {
	if v.Field > v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not greater than %d", v.Field, v.ComparedField))
}
