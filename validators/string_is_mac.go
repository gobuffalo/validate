package validators

import (
	"fmt"
	"net"

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
