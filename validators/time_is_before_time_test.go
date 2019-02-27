package validators

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_TimeIsBeforeTime(t *testing.T) {

	r := require.New(t)

	now := time.Now()
	v := TimeIsBeforeTime{
		Name: "OpensAt", Field: now,
		ComparedName: "ClosesAt", ComparedField: now.Add(100000),
	}

	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v.ComparedField = now.Add(-100000)
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"OpensAt must be before ClosesAt"}, e.Get("OpensAt"))
}
