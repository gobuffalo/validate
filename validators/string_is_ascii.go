package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsASCII is a validator object
type StringIsASCII struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains anything except for ASCII characters.
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

	e.Add(v.Name, fmt.Sprintf("%s must contain ASCII chars only", v.Name))
}

// SetField sets validator field.
func (v *StringIsASCII) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsASCII) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
