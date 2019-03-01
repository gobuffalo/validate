package validators

import (
	"fmt"
	"regexp"

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

// SetField sets validator field.
func (v *StringIsBase64) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsBase64) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
