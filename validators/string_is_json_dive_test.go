package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsJSONDive(t *testing.T) {

	r := require.New(t)

	field := []string{"{\"test\": \"sure\"}", "   {\"test\": \"sure\"}    ", "123"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsJSON{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{" ", "abc", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsJSON{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())
}
