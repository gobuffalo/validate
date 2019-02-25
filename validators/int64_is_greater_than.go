package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64IsGreaterThan is a validator object
type Int64IsGreaterThan struct {
	Name     string
	Field    int64
	Compared int64
	Message  string
}

// Validate adds an error if the field is not greater than the compared value.
func (v *Int64IsGreaterThan) Validate(e *validator.Errors) {
	if v.Field > v.Compared {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not greater than %d", v.Field, v.Compared))
}
