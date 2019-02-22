package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreNotEqual is a validator object
type IntsAreNotEqual struct {
	ValueOne int
	ValueTwo int
	Name     string
	Message  string
}

// Validate is a validator that compares two integers and will add an error if they are equal
func (v *IntsAreNotEqual) Validate(e *validator.Errors) {
	if v.ValueOne == v.ValueTwo {
		e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.ValueOne, v.ValueTwo))
	}
}
