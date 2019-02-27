package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsHexadecimal is a validator object.
type StringIsHexadecimal struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not in a hexadecimal format.
func (v *StringIsHexadecimal) Validate(e *validator.Errors) {

	if rxHexadecimal.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be a hexadecimal number", v.Name))
}
