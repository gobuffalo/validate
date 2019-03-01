package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsBase64Dive(t *testing.T) {

	r := require.New(t)

	field := []string{"xoP4nZV8Gv9ceg==", "n/GNBg=="} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsBase64{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"xoP4nZV8Gv9ceg==", "n/GNBg==", "xoP4nZ V8Gv9ceg==", " n/GNBg==", " ", ""} // 4 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsBase64{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
