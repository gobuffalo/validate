package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsBase64 is a validator object
type StringIsBase64 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field is not base64 encoded.
func (v *StringIsBase64) Validate(e *validator.Errors) {

	// base64 is valid
	if rxBase64.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be base64 encoded", v.Name))
}
