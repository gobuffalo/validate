package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPortDive(t *testing.T) {

	r := require.New(t)

	field := []string{"1", "123", "65535"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsPort{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"1", "123", "65535", "1 ", "65536", " ", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsPort{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
