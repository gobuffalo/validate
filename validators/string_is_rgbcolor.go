package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsRGBcolor is a validator object
type StringIsRGBcolor struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is a RGB color in form rgb(RRR, GGG, BBB).
func (v *StringIsRGBcolor) Validate(e *validator.Errors) {

	if rxRGBcolor.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be a RGB color in form rgb(RRR, GGG, BBB)", v.Name))
}
