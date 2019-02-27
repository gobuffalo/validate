package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// Int64sAreEqual is a validator object.
type Int64sAreEqual struct {
	Name     string
	Field    int64
	Compared int64
}

// Validate add an error if the field is not equal to the compared value.
func (v *Int64sAreEqual) Validate(e *validator.Errors) {
	if v.Field == v.Compared {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%d is not equal to %d", v.Field, v.Compared))
}
