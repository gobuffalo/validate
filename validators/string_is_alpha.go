package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsAlpha is a validator object.
type StringIsAlpha struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains anything except for latin letters.
// Empty string is valid.
func (v *StringIsAlpha) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alpha is valid, not trimming spaces
	if rxAlpha.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must contain only letters (a-zA-Z)", v.Name))
}

// SetField sets validator field.
func (v *StringIsAlpha) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsAlpha) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
