package validators

import (
	"fmt"
	"strconv"

	"github.com/s3rj1k/validator"
)

// StringIsPort is a validator object
type StringIsPort struct {
	Name    string
	Field   string
	Message string
}

// Validate if a string represents a valid port
func (v *StringIsPort) Validate(e *validator.Errors) {

	if i, err := strconv.Atoi(v.Field); err == nil && i > 0 && i < 65536 {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must represent a valid port", v.Name))
}
