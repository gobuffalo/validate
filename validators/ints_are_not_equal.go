package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreNotEqual is a validator object.
type IntsAreNotEqual struct {
	Name     string
	Field    int
	Compared int
}

// Validate add an error if the field is equal to the compared value.
func (v *IntsAreNotEqual) Validate(e *validator.Errors) {
	if v.Field != v.Compared {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.Compared))
}
