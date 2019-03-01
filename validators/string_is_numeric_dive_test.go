package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNumericDive(t *testing.T) {

	r := require.New(t)

	field := []string{"123", "555", "123 ", "55 55", "", "  "} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsNumeric{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())

	field = []string{"ab", "123#", "1 ", "11 11", "", "  "} // 5 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsNumeric{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(5, e.Count())
}
