package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64sAreEqual is a validator object
type Int64sAreEqual struct {
	ValueOne int64
	ValueTwo int64
	Name     string
	Message  string
}

// Validate is a validator that will compare two integers and add an error if they are not equal
func (v *Int64sAreEqual) Validate(e *validator.Errors) {
	if v.ValueOne != v.ValueTwo {
		e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.ValueOne, v.ValueTwo))
	}
}
