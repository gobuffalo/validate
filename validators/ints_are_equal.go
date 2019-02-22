package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreEqual is a validator object
type IntsAreEqual struct {
	ValueOne int
	ValueTwo int
	Name     string
	Message  string
}

// Validate is a validator that will compare two integers and add an error if they are not equal
func (v *IntsAreEqual) Validate(e *validator.Errors) {
	if v.ValueOne != v.ValueTwo {
		e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.ValueOne, v.ValueTwo))
	}
}
