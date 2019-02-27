package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsPrintableASCII is a validator object.
type StringIsPrintableASCII struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains anything except for printable ASCII characters.
// Empty string is valid.
func (v *StringIsPrintableASCII) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// ASCII is valid
	if rxPrintableASCII.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must contain printable ASCII chars only", v.Name))
}
