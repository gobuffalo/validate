package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntIsPresent is a validator object.
type IntIsPresent struct {
	Name    string
	Field   int
	Message string
}

// Validate adds an error if the field equals 0.
func (v *IntIsPresent) Validate(e *validator.Errors) {
	if v.Field != 0 {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
