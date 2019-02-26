package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreNotEqual is a validator object
type IntsAreNotEqual struct {
	Name     string
	Field    int
	Compared int
	Message  string
}

// Validate is a validator that compares two integers and will add an error if they are equal
func (v *IntsAreNotEqual) Validate(e *validator.Errors) {
	if v.Field == v.Compared {
		e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.Compared))
	}
}
