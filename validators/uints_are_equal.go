package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintsAreEqual is a validator object
type UintsAreEqual struct {
	Name          string
	Field         uint
	ComparedName  string
	ComparedField uint
}

// Validate adds an error if the Field is not equal to the ComparedField.
func (v *UintsAreEqual) Validate(e *validator.Errors) {
	if v.Field == v.ComparedField {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.ComparedField))
}
