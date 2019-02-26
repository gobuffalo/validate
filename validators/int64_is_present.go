package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64IsPresent is a validator object.
type Int64IsPresent struct {
	Name    string
	Field   int64
	Message string
}

// Validate adds an error if the field equals 0.
func (v *Int64IsPresent) Validate(e *validator.Errors) {
	if v.Field != 0 {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
