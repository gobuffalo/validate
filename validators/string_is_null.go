package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNull is a validator object.
type StringIsNull struct {
	Name    string
	Field   string
	Message string
}

// isNullString is wrapper func
func isNullString(str string) bool {
	if len(str) == 0 { // nolint: megacheck
		return true
	}

	return false
}

// Validate adds an error if the field is not empty.
func (v *StringIsNull) Validate(e *validator.Errors) {
	if isNullString(v.Field) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be empty", v.Name))
}
