package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNumeric is a validator object
type StringIsNumeric struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string contains only numbers. Empty string is valid.
func (v *StringIsNumeric) Validate(e *validator.Errors) {

	// null string is valid
	if IsNull(v.Field) {
		return
	}

	// alphanum is valid, not trimming spaces
	if rxNumeric.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain only numbers", v.Name))
}
