package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPresentDive(t *testing.T) {

	r := require.New(t)

	field := []string{"abc", " a", "3  3"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsPresent{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"", "   ", "1", "a"}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsPresent{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(2, e.Count())
}
