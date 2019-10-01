package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

// FloatIsLessThan is a validator that will compare two floats and add
// an error if the field is not less than the compared value
type FloatIsLessThan struct {
	Name     string
	Field    float64
	Compared float64
	Message  string
}

// IsValid adds an error if the field is not less than the compared value.
func (v *FloatIsLessThan) IsValid(errors *validate.Errors) {
	if v.Field < v.Compared {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%f is not less than %f.", v.Field, v.Compared))
}
