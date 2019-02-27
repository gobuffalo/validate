package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntsAreEqual is a validator object.
type IntsAreEqual struct {
	Name     string
	Field    int
	Compared int
}

// Validate adds an error if the field is not equal to the compared value.
func (v *IntsAreEqual) Validate(e *validator.Errors) {
	if v.Field == v.Compared {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.Compared))
}
