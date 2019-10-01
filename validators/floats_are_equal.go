package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

// FloatsAreEqual is a validator that will compare two floats and add
// an error if they are not equal
type FloatsAreEqual struct {
	ValueOne float64
	ValueTwo float64
	Name     string
	Message  string
}

// IsValid adds an error if they are not equal.
func (v *FloatsAreEqual) IsValid(errors *validate.Errors) {
	if v.ValueOne != v.ValueTwo {
		errors.Add(GenerateKey(v.Name), fmt.Sprintf("%f is not equal to %f", v.ValueOne, v.ValueTwo))
	}
}
