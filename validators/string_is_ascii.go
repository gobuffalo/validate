package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsASCII is a validator object
type StringIsASCII struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field contains anything except for ASCII characters.
// Empty string is valid.
func (v *StringIsASCII) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// ASCII is valid
	if rxASCII.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain ASCII chars only", v.Name))
}
