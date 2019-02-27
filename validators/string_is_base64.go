package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsBase64 is a validator object
type StringIsBase64 struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not base64 encoded.
func (v *StringIsBase64) Validate(e *validator.Errors) {

	// base64 is valid
	if rxBase64.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be base64 encoded", v.Name))
}
