package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64sAreNotEqual is a validator object
type Int64sAreNotEqual struct {
	ValueOne int64
	ValueTwo int64
	Name     string
	Message  string
}

// Validate is a validator that compares two integers and will add an error if they are equal
func (v *Int64sAreNotEqual) Validate(e *validator.Errors) {
	if v.ValueOne == v.ValueTwo {
		e.Add(v.Name, fmt.Sprintf("%d is equal to %d", v.ValueOne, v.ValueTwo))
	}
}
