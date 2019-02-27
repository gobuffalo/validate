package validators

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_TimeIsAfterTime(t *testing.T) {

	r := require.New(t)

	now := time.Now()
	v := TimeIsAfterTime{
		Name: "OpensAt", Field: now.Add(100000),
		ComparedName: "Now", ComparedField: now,
	}

	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v.ComparedField = now.Add(200000)
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"OpensAt must be after Now"}, e.Get("OpensAt"))
}
