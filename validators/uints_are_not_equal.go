package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintsAreNotEqual is a validator object
type UintsAreNotEqual struct {
	Name          string
	Field         uint
	ComparedName  string
	ComparedField uint
}

// Validate adds an error if the Field is equal to the ComparedField.
func (v *UintsAreNotEqual) Validate(e *validator.Errors) {
	if v.Field != v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.ComparedField))
}
