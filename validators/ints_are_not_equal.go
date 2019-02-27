package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreNotEqual is a validator object.
type IntsAreNotEqual struct {
	Name          string
	Field         int
	ComparedName  string
	ComparedField int
}

// Validate add an error if the Field is equal to the ComparedField.
func (v *IntsAreNotEqual) Validate(e *validator.Errors) {
	if v.Field != v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.ComparedField))
}
