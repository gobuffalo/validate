package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

type IntIsGreaterThanOrEqualTo struct {
	Name     string
	Field    int
	Compared int
	Message  string
}

// IsValid adds an error if the field is not greater than the compared value.
func (v *IntIsGreaterThanOrEqualTo) IsValid(errors *validate.Errors) {
	if v.Field >= v.Compared {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%d is not greater than or equal to %d.", v.Field, v.Compared))
}
