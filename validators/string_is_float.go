package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsFloat is a validator object
type StringIsFloat struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is a float. Empty string is valid.
func (v *StringIsFloat) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// Int is valid
	if rxFloat.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be a float", v.Name))
}
