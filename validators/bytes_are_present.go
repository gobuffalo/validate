package validators

import (
	"fmt"

	"github.com/gobuffalo/validate/v3"
)

// BytesArePresent is a validator that will adds an error if the array of bytes is empty
type BytesArePresent struct {
	Name    string
	Field   []byte
	Message string
}

// IsValid adds an error if the array of bytes is empty.
func (v *BytesArePresent) IsValid(errors *validate.Errors) {
	if len(v.Field) > 0 {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%s can not be blank.", v.Name))
}
