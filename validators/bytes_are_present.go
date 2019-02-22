package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// BytesArePresent is a validator object
type BytesArePresent struct {
	Name    string
	Field   []byte
	Message string
}

// Validate adds an error if the field is not empty.
func (v *BytesArePresent) Validate(e *validator.Errors) {
	if len(v.Field) > 0 {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank.", v.Name))
}
