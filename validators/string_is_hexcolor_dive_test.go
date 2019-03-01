package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsHexcolorDive(t *testing.T) {

	r := require.New(t)

	field := []string{"#1f1f1F", "#AFAFAF", "#1AFFa1"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsHexcolor{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"#1AFFa1", "#1f1f1F", "#1f1f1F ", "#afaf", "#F0h", "", "  "}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsHexcolor{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(5, e.Count())
}
