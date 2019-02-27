package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsFloat is a validator object.
type StringIsFloat struct {
	Name  string
	Field string
}

// Validate add an error if the Field is not a float. Empty string is valid.
func (v *StringIsFloat) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// float is valid
	if rxFloat.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be a float", v.Name))
}
