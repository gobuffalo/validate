package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64ArrayIsPresent is a validator object.
type Int64ArrayIsPresent struct {
	Name  string
	Field []int64
}

// Validate adds an error if the field is an empty array.
func (v *Int64ArrayIsPresent) Validate(e *validator.Errors) {
	if len(v.Field) > 0 {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be empty", v.Name))
}
