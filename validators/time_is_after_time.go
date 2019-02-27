package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeIsAfterTime is a validator object.
type TimeIsAfterTime struct {
	Name          string
	Field         time.Time
	ComparedName  string
	ComparedField time.Time
}

// Validate adds an error if the Field time is not after the ComparedField time.
func (v *TimeIsAfterTime) Validate(e *validator.Errors) {
	if v.Field.UnixNano() > v.ComparedField.UnixNano() {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be after %s", v.Name, v.ComparedName))
}
