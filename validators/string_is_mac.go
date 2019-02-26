package validators

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsMAC is a validator object.
type StringIsMAC struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field is not formatted as a MAC address.
func (v *StringIsMAC) Validate(e *validator.Errors) {

	// using net ParseMAC
	_, err := net.ParseMAC(v.Field)
	if err == nil {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be valid MAC address", v.Name))
}
