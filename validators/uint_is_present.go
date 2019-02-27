package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintIsPresent is a validator object
type UintIsPresent struct {
	Name  string
	Field uint
}

// Validate adds an error if the Field is equal to 0.
func (v *UintIsPresent) Validate(e *validator.Errors) {
	if v.Field != 0 {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
