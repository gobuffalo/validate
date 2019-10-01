package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

// FloatIsGreaterThanOrEqualTo is a validator that will compare two floats and add
// an error if the field is not greater than or equal the compared value
type FloatIsGreaterThanOrEqualTo struct {
	Name     string
	Field    float64
	Compared float64
	Message  string
}

// IsValid adds an error if the field is not greater than or equal the compared value.
func (v *FloatIsGreaterThanOrEqualTo) IsValid(errors *validate.Errors) {
	if v.Field >= v.Compared {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%f is not greater than or equal to %f.", v.Field, v.Compared))
}
