package validators

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsAbsPath is a validator object
type StringIsAbsPath struct {
	Name  string
	Field string
}

// Validate adds an error if Field is not an absolute path.
func (v *StringIsAbsPath) Validate(e *validator.Errors) {
	if v.Field == filepath.Clean(v.Field) && filepath.IsAbs(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must be absolute", v.Field))
}

// SetField sets validator field.
func (v *StringIsAbsPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsAbsPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
