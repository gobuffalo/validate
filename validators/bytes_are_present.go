package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// BytesArePresent is a validator object.
type BytesArePresent struct {
	Name  string
	Field []byte
}

// Validate adds an error if the field is empty.
func (v *BytesArePresent) Validate(e *validator.Errors) {
	if len(v.Field) > 0 {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
