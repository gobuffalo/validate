package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNullDive(t *testing.T) {

	r := require.New(t)

	field := []string{"", "", ""}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsNull{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"", " 12.5", "   ", "a", " ", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsNull{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
