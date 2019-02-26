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
	Message  string
}

// Validate is a validator that compares two unsigned integers and will add an error if they are equal
func (v *UintsAreNotEqual) Validate(e *validator.Errors) {
	if v.Field == v.Compared {
		e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.Field, v.Compared))
	}
}
