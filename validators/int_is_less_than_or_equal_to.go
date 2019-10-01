package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

// IntIsLessThanOrEqualTo is a validator that will compare two ints and add
// an error if the field is not less than or equal the compared value
type IntIsLessThanOrEqualTo struct {
	Name     string
	Field    int
	Compared int
	Message  string
}

// IsValid adds an error if the field is not less than or equal the compared value.
func (v *IntIsLessThanOrEqualTo) IsValid(errors *validate.Errors) {
	if v.Field <= v.Compared {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%d is not less than or equal to %d.", v.Field, v.Compared))
}
