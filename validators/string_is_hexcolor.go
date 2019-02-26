package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsHexcolor is a validator object.
type StringIsHexcolor struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field is not formatted as a hexadecimal color.
// Leading '#' is required (e.g. "#1f1f1F", "#F00").
func (v *StringIsHexcolor) Validate(e *validator.Errors) {

	if rxHexcolor.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be a hexadecimal color", v.Name))
}
