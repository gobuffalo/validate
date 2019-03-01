package validators

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/s3rj1k/validator"
)

// StringIsPort is a validator object.
type StringIsPort struct {
	Name  string
	Field string
}

// Validate adds an error if the Field does not represent a valid port.
func (v *StringIsPort) Validate(e *validator.Errors) {

	if i, err := strconv.Atoi(v.Field); err == nil && i > 0 && i < 65536 {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must represent a valid port", v.Name))
}

// SetField sets validator field.
func (v *StringIsPort) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPort) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
