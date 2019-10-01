package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

type FloatIsGreaterThan struct {
	Name     string
	Field    float64
	Compared float64
	Message  string
}

// IsValid adds an error if the field is not greater than the compared value.
func (v *FloatIsGreaterThan) IsValid(errors *validate.Errors) {
	if v.Field > v.Compared {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%f is not greater than %f.", v.Field, v.Compared))
}
