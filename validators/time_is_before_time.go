package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeIsBeforeTime is a validator object.
type TimeIsBeforeTime struct {
	FirstName  string
	FirstTime  time.Time
	SecondName string
	SecondTime time.Time
}

// Validate adds an error if the FirstTime is after the SecondTime.
func (v *TimeIsBeforeTime) Validate(e *validator.Errors) {
	if v.FirstTime.UnixNano() <= v.SecondTime.UnixNano() {
		return
	}

	e.Add(v.FirstName, fmt.Sprintf("%s must be before %s", v.FirstName, v.SecondName))
}
