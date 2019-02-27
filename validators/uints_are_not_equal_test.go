package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_UintsAreNotEqual(t *testing.T) {

	r := require.New(t)

	v := UintsAreNotEqual{Name: "Number", Field: 2, ComparedField: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = UintsAreNotEqual{Name: "Number", Field: 2, ComparedField: 2}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"2 is equal to 2"}, e.Get("Number"))
}
