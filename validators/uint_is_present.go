package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintIsPresent is a validator object
type UintIsPresent struct {
	Name    string
	Field   uint
	Message string
}

// Validate adds an error if the field equals 0.
func (v *UintIsPresent) Validate(e *validator.Errors) {
	if v.Field != 0 {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
