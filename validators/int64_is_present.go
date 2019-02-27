package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64IsPresent is a validator object.
type Int64IsPresent struct {
	Name  string
	Field int64
}

// Validate adds an error if the Field equals to 0.
func (v *Int64IsPresent) Validate(e *validator.Errors) {
	if v.Field != 0 {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
