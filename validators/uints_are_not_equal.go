package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintsAreNotEqual is a validator object
type UintsAreNotEqual struct {
	Name     string
	Field    uint
	Compared uint
}

// Validate adds an error if the field is equal to the compared value.
func (v *UintsAreNotEqual) Validate(e *validator.Errors) {
	if v.Field != v.Compared {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.Compared))
}
