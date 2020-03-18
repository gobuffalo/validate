package validators

import (
	"testing"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/stretchr/testify/require"
)

func Test_TimeIsBeforeTime(t *testing.T) {
	r := require.New(t)
	now := time.Now()
	v := TimeIsBeforeTime{
		FirstName: "Opens At", FirstTime: now,
		SecondName: "Closes At", SecondTime: now.Add(100000),
	}

	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(0, errors.Count())

	v.SecondTime = now.Add(-100000)
	v.IsValid(errors)

	r.Equal(1, errors.Count())
	r.Equal(errors.Get("opens_at"), []string{"Opens At must be before Closes At."})

	errors = validate.NewErrors()
	v.Message = "OpensAt must be earlier than ClosesAt."

	v.IsValid(errors)

	r.Equal(1, errors.Count())
	r.Equal(errors.Get("opens_at"), []string{"OpensAt must be earlier than ClosesAt."})

	firstTime := now.AddDate(-400, 0, 0)
	v = TimeIsBeforeTime{
		FirstName: "Opens At", FirstTime: firstTime,
		SecondName: "Now", SecondTime: now,
	}

	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(0, errors.Count())
}
