package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeAfterTime is a validator object
type TimeAfterTime struct {
	FirstName  string
	FirstTime  time.Time
	SecondName string
	SecondTime time.Time
	Message    string
}

// Validate adds an error if the FirstTime is not after the SecondTime.
func (v *TimeAfterTime) Validate(e *validator.Errors) {
	if v.FirstTime.UnixNano() >= v.SecondTime.UnixNano() {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.FirstName, v.Message)
		return
	}

	e.Add(v.FirstName, fmt.Sprintf("%s must be after %s", v.FirstName, v.SecondName))
}
