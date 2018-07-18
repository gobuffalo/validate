package validators

import (
	"fmt"

	"github.com/gobuffalo/validate"
)

type BytesArePresent struct {
	Name    string
	Field   []byte
	Message string
}

// IsValid adds an error if the field is not empty.
func (v *BytesArePresent) IsValid(errors *validate.Errors) {
	if len(v.Field) > 0 {
		return
	}

<<<<<<< HEAD
	if len(v.Message) > 0 {
=======
	if len(v.Message) >= 0 {
>>>>>>> 5759c934d8b995425f342bae010a3f3dd393ea2a
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateKey(v.Name), fmt.Sprintf("%s can not be blank.", v.Name))
}
