package validators

import (
	"fmt"
	"time"

	"github.com/markbates/validate"
)

type TimeIsAfterTime struct {
	FirstName  string
	FirstTime  time.Time
	SecondName string
	SecondTime time.Time
}

func (v *TimeIsAfterTime) IsValid(errors *validate.Errors) {
	if v.FirstTime.UnixNano() < v.SecondTime.UnixNano() {
		errors.Add(GenerateKey(v.FirstName), fmt.Sprintf("%s must be after %s.", v.FirstName, v.SecondName))
	}
}
