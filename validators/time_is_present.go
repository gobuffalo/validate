package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeIsPresent is a validator object
type TimeIsPresent struct {
	Name    string
	Field   time.Time
	Message string
}

// Validate adds an error if the field is not a valid time.
func (v *TimeIsPresent) Validate(e *validator.Errors) {
	t := time.Time{}
	if v.Field.UnixNano() != t.UnixNano() {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank.", v.Name))
}
