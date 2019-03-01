package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsAlphaDive(t *testing.T) {

	r := require.New(t)

	field := []string{"Need", "onlY", "Letters", "No", "WhiteSpaces", ""} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsAlpha{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"These", "HAV#", "s0m3", "numb3rs", "", "al so"} // 4 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsAlpha{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
