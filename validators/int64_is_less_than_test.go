package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_Int64IsLessThan(t *testing.T) {

	r := require.New(t)

	v := Int64IsLessThan{Name: "Number", Field: 1, ComparedField: 2}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = Int64IsLessThan{Name: "Number", Field: 1, ComparedField: 0}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"1 is not less than 0"}, e.Get("Number"))

	v = Int64IsLessThan{Name: "Number"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"0 is not less than 0"}, e.Get("Number"))
}
