package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsInt is a validator object.
type StringIsInt struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not an integer.
// Leading sign is allowed. Empty string is valid.
func (v *StringIsInt) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// Int is valid
	if rxInt.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be an integer", v.Name))
}
