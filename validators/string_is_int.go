package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsInt is a validator object
type StringIsInt struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is an integer. Empty string is valid.
func (v *StringIsInt) Validate(e *validator.Errors) {

	// null string is valid
	if IsNull(v.Field) {
		return
	}

	// Int is valid
	if rxInt.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be an integer", v.Name))
}
