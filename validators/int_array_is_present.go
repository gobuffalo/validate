package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntArrayIsPresent is a validator object.
type IntArrayIsPresent struct {
	Name    string
	Field   []int
	Message string
}

// Validate adds an error if the field is an empty array.
func (v *IntArrayIsPresent) Validate(e *validator.Errors) {
	if len(v.Field) > 0 {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be empty", v.Name))
}
