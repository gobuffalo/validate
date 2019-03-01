package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringHasUpperCase is a validator object.
type StringHasUpperCase struct {
	Name  string
	Field string
}

// Validate adds an error if the Field has not uppercased letters. Empty string is valid.
func (v *StringHasUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if rxHasUpperCase.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must contain at least 1 uppercase", v.Name))
}

// SetField sets validator field.
func (v *StringHasUpperCase) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasUpperCase) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
