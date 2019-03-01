package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsAlphaNumDive(t *testing.T) {

	r := require.New(t)

	field := []string{"Need", "0nlY", "Letters", "And", "Numb3r5", "", "No", "whitespac3s"} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsAlphaNum{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"These", "HAV#", "s0m3", "oth&r", "sym bols", " also", ""} // 4 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsAlphaNum{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
