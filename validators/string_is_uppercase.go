package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsUpperCase is a validator object.
type StringIsUpperCase struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not uppercased. Empty string is valid.
func (v *StringIsUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if v.Field == strings.ToUpper(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be uppercase", v.Name))
}

// SetField sets validator field.
func (v *StringIsUpperCase) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUpperCase) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
