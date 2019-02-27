package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreEqual is a validator object.
type IntsAreEqual struct {
	Name          string
	Field         int
	ComparedName  string
	ComparedField int
}

// Validate adds an error if the Field is not equal to the ComparedField.
func (v *IntsAreEqual) Validate(e *validator.Errors) {
	if v.Field == v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.ComparedField))
}
