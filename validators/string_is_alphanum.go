package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsAlphaNum is a validator object.
type StringIsAlphaNum struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field contains symbols except for arabic numerals and latin letters.
// Empty string is valid.
func (v *StringIsAlphaNum) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alphanum is valid, not trimming spaces
	if rxAlphanumeric.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain only numbers and/or letters", v.Name))
}
