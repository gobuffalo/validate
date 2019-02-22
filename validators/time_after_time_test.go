package validators

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_TimeAfterTime(t *testing.T) {

	r := require.New(t)

	now := time.Now()
	v := TimeAfterTime{
		FirstName: "OpensAt", FirstTime: now.Add(100000),
		SecondName: "Now", SecondTime: now,
	}

	e := validator.NewErrors()
	v.Validate(e)

	r.Equal(0, e.Count())

	v.SecondTime = now.Add(200000)
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"OpensAt must be after Now."}, e.Get("OpensAt"))

	e = validator.NewErrors()
	v.Message = "OpensAt must be later than Now."

	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"OpensAt must be later than Now."}, e.Get("OpensAt"))
}
