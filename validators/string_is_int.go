package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsInt is a validator object.
type StringIsInt struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not an integer.
// Leading sign is allowed. Empty string is valid.
func (v *StringIsInt) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// Int is valid
	if rxInt.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be an integer", v.Name))
}

// SetField sets validator field.
func (v *StringIsInt) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsInt) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
