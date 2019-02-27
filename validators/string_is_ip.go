package validators

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsIP is a validator object.
type StringIsIP struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a valid IP address version 4 or 6.
func (v *StringIsIP) Validate(e *validator.Errors) {

	// ParseIP has to return a non-nil value
	if net.ParseIP(v.Field) != nil {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be either IP version 4 or 6", v.Name))
}
