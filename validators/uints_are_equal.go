package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// UintsAreEqual is a validator object
type UintsAreEqual struct {
	Name     string
	Field    uint
	Compared uint
}

// Validate adds an error if the field is not equal to the compared value.
func (v *UintsAreEqual) Validate(e *validator.Errors) {
	if v.Field == v.Compared {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.Compared))
}
