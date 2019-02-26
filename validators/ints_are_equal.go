package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreEqual is a validator object
type IntsAreEqual struct {
	Name     string
	Field    int
	Compared int
	Message  string
}

// Validate is a validator that will compare two integers and add an error if they are not equal
func (v *IntsAreEqual) Validate(e *validator.Errors) {
	if v.Field != v.Compared {
		e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.Compared))
	}
}
