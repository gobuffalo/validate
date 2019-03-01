package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasNoPrefix is a validator object.
type StringHasNoPrefix struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
}

// Validate adds an error if the Field is prefixed with ComparedField.
func (v *StringHasNoPrefix) Validate(e *validator.Errors) {

	if !strings.HasPrefix(v.Field, v.ComparedField) {
		return
	}

	if len(v.ComparedName) == 0 {
		e.Add(v.Name, fmt.Sprintf("'%s' starts with '%s'", v.Field, v.ComparedField))
	} else {
		e.Add(v.Name, fmt.Sprintf("'%s' starts with content of '%s'", v.Name, v.ComparedName))
	}
}

// SetField sets validator field.
func (v *StringHasNoPrefix) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasNoPrefix) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
