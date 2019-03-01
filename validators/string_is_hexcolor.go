package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsHexcolor is a validator object.
type StringIsHexcolor struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not formatted as a hexadecimal color.
// Leading '#' is required (e.g. "#1f1f1F", "#F00").
func (v *StringIsHexcolor) Validate(e *validator.Errors) {

	if rxHexcolor.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be a hexadecimal color", v.Name))
}

// SetField sets validator field.
func (v *StringIsHexcolor) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsHexcolor) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
