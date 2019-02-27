package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_UintIsGreaterThan(t *testing.T) {

	r := require.New(t)

	v := UintIsGreaterThan{Name: "Number", Field: 2, ComparedField: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = UintIsGreaterThan{Name: "Number", Field: 1, ComparedField: 2}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"1 is not greater than 2"}, e.Get("Number"))

	v = UintIsGreaterThan{"Number", 0, "Number2", 0}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"0 is not greater than 0"}, e.Get("Number"))
}
