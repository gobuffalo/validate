package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsMAC is a validator object.
type StringIsMAC struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a MAC address.
func (v *StringIsMAC) Validate(e *validator.Errors) {

	// using net ParseMAC
	_, err := net.ParseMAC(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be valid MAC address", v.Name))
}

// SetField sets validator field.
func (v *StringIsMAC) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsMAC) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
