package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsHexadecimal is a validator object
type StringIsHexadecimal struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is a hexadecimal number.
func (v *StringIsHexadecimal) Validate(e *validator.Errors) {

	if rxHexadecimal.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be a hexadecimal number", v.Name))
}
