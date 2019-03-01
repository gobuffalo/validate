package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsRGBColorDive(t *testing.T) {

	r := require.New(t)

	field := []string{"rgb(0,0,0)", "rgb(255,255,255)"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsRGBcolor{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"rgb(0,0,0)", "rgb(255,255,255)", "RGB(255,255,255)", "RGB(256,255,255)", "RGB(,255,255)", " ", "", "rgb(0,0,0) "}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsRGBcolor{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(6, e.Count())
}
