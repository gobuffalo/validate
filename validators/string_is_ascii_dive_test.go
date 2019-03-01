package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsASCIIDive(t *testing.T) {

	r := require.New(t)

	field := []string{"Need", "0nlY", " ASC11", "!$#%()-=<>etc...,@", " ", ""} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsASCII{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"тут", "ошибочка", " s0m3", "oth &r", "sym bols", " ", ""} // 2 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsASCII{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(2, e.Count())
}
