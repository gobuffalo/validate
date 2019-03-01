package validators

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/s3rj1k/validator"
)

// StringLengthInRange is a validator object.
type StringLengthInRange struct {
	Name  string
	Field string
	Min   int
	Max   int
}

// Validate adds an error if the Field length is not in range between Min and Max (inclusive).
// It is possible to provide either both or one of the Min/Max values.
func (v *StringLengthInRange) Validate(e *validator.Errors) {
	strLength := utf8.RuneCountInString(v.Field)

	min := v.Min
	max := v.Max

	if max == 0 {
		max = strLength
	}

	if strLength >= min && strLength <= max {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s not in range(%d, %d)", v.Name, min, max))
}

// SetField sets validator field.
func (v *StringLengthInRange) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringLengthInRange) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
