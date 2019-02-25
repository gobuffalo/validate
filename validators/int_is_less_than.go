package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntIsLessThan is a validator object
type IntIsLessThan struct {
	Name     string
	Field    int
	Compared int
	Message  string
}

// Validate adds an error if the field is not less than the compared value.
func (v *IntIsLessThan) Validate(e *validator.Errors) {
	if v.Field < v.Compared {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not less than %d", v.Field, v.Compared))
}
