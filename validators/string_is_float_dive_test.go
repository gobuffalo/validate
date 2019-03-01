package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsFloatDive(t *testing.T) {

	r := require.New(t)

	field := []string{"1.5", "+1", "0", "-12.5"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsFloat{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"11", "12.5", "14,5", "12. 1", " 0", " ", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsFloat{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
