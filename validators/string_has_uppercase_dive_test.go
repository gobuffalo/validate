package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasUpperCaseDive(t *testing.T) {

	r := require.New(t)

	field := []string{"Need", "AT", "LeAST", "One", "Upper Case", ""} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringHasUpperCase{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"These", "HAVE", "but", "these", "do not", ""} // 3 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringHasUpperCase{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())
}
