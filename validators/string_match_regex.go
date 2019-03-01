package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringMatchRegex is a validator object.
type StringMatchRegex struct {
	Name  string
	Field string
	Regex string
}

// Validate adds an error if the Field does not match regular expression Regex.
func (v *StringMatchRegex) Validate(e *validator.Errors) {
	r := regexp.MustCompile(v.Regex)
	if r.Match([]byte(v.Field)) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s does not match the expected format", v.Name))
}

// SetField sets validator field.
func (v *StringMatchRegex) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringMatchRegex) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
