package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIntDive(t *testing.T) {

	r := require.New(t)

	field := []string{"1", "+1", "0", "-12"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsInt{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"11", "12.5", "a", " 11", "1 1", " ", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsInt{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(5, e.Count())
}
