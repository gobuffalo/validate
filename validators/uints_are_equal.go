package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintsAreEqual is a validator object
type UintsAreEqual struct {
	Name     string
	Field    uint
	Compared uint
	Message  string
}

// Validate is a validator that will compare two unsigned integers and add an error if they are not equal
func (v *UintsAreEqual) Validate(e *validator.Errors) {
	if v.Field != v.Compared {
		e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.Compared))
	}
}
