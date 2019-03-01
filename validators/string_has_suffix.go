package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasSuffix is a validator object.
type StringHasSuffix struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
}

// Validate adds an error if the Field is not suffixed with ComparedField.
func (v *StringHasSuffix) Validate(e *validator.Errors) {

	if strings.HasSuffix(v.Field, v.ComparedField) {
		return
	}

	if len(v.ComparedName) == 0 {
		e.Add(v.Name, fmt.Sprintf("'%s' does not end with '%s'", v.Field, v.ComparedField))
	} else {
		e.Add(v.Name, fmt.Sprintf("'%s' does not end with content of '%s'", v.Name, v.ComparedName))
	}
}

// SetField sets validator field.
func (v *StringHasSuffix) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasSuffix) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
