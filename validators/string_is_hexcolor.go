package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsHexcolor is a validator object.
type StringIsHexcolor struct {
	Name  string
	Field string
}

// Validate adds an error if the field is not formatted as a hexadecimal color.
// Leading '#' is required (e.g. "#1f1f1F", "#F00").
func (v *StringIsHexcolor) Validate(e *validator.Errors) {

	if rxHexcolor.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be a hexadecimal color", v.Name))
}
