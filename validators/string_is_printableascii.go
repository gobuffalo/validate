package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsPrintableASCII is a validator object
type StringIsPrintableASCII struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string contains printable ASCII chars only. Empty string is valid.
func (v *StringIsPrintableASCII) Validate(e *validator.Errors) {

	// null string is valid
	if IsNull(v.Field) {
		return
	}

	// ASCII is valid
	if rxPrintableASCII.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain printable ASCII chars only", v.Name))
}
