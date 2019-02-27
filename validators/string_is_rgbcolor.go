package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsRGBcolor is a validator object.
type StringIsRGBcolor struct {
	Name  string
	Field string
}

// Validate adds an error if the field is not formatted as an RGB color.
// Expected format is "rgb(RRR, GGG, BBB)".
func (v *StringIsRGBcolor) Validate(e *validator.Errors) {

	if rxRGBcolor.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be a RGB color in form rgb(RRR, GGG, BBB)", v.Name))
}
