package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeIsPresent is a validator object.
type TimeIsPresent struct {
	Name  string
	Field time.Time
}

// Validate adds an error if the Field is the time default value.
func (v *TimeIsPresent) Validate(e *validator.Errors) {
	t := time.Time{}
	if v.Field.UnixNano() != t.UnixNano() {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
