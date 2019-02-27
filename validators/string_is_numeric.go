package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNumeric is a validator object.
type StringIsNumeric struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not numeric. Empty string is valid.
func (v *StringIsNumeric) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alphanum is valid, not trimming spaces
	if rxNumeric.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must contain only numbers", v.Name))
}
