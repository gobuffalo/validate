package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntIsPresent is a validator object.
type IntIsPresent struct {
	Name  string
	Field int
}

// Validate adds an error if the field equals 0.
func (v *IntIsPresent) Validate(e *validator.Errors) {
	if v.Field != 0 {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
