package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsAlpha is a validator object
type StringIsAlpha struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string contains only letters (a-zA-Z). Empty string is valid.
func (v *StringIsAlpha) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alpha is valid, not trimming spaces
	if rxAlpha.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain only letters (a-zA-Z)", v.Name))
}
