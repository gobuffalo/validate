package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

// FloatsAreNotEqual is a validator that compares two floats and will add
// an error if they are equal
type FloatsAreNotEqual struct {
	ValueOne float64
	ValueTwo float64
	Name     string
	Message  string
}

func (v *FloatsAreNotEqual) IsValid(errors *validate.Errors) {
	if v.ValueOne == v.ValueTwo {
		errors.Add(GenerateKey(v.Name), fmt.Sprintf("%f is equal to %f", v.ValueOne, v.ValueTwo))
	}
}
