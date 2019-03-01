package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUpperCaseDive(t *testing.T) {

	r := require.New(t)

	field := []string{"NEED", "ONLY", "UPPER CASES", "  ", ""}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsUpperCase{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"These", "HAVE", "some", "Lowers", "", " "}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsUpperCase{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())
}
