package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsAlphaNum is a validator object.
type StringIsAlphaNum struct {
	Name  string
	Field string
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

	e.Add(v.Name, fmt.Sprintf("%s must contain only numbers and/or letters", v.Name))
}
