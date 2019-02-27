package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IntIsGreaterThan(t *testing.T) {

	r := require.New(t)

	v := IntIsGreaterThan{Name: "Number", Field: 2, ComparedField: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = IntIsGreaterThan{Name: "Number", Field: 1, ComparedField: 2}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"1 is not greater than 2"}, e.Get("Number"))

	v = IntIsGreaterThan{Name: "Number", Field: 1, ComparedField: 1}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"1 is not greater than 1"}, e.Get("Number"))
}
