package validators

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsIP is a validator object
type StringIsIP struct {
	Name    string
	Field   string
	Message string
}

// Validate if a string is either IP version 4 or 6.
func (v *StringIsIP) Validate(e *validator.Errors) {

	// ParseIP has to return a non-nil value
	if net.ParseIP(v.Field) != nil {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be either IP version 4 or 6", v.Name))
}
