package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringExclusionDive(t *testing.T) {

	r := require.New(t)

	blackl := []string{}
	field := []string{"This", "is", "good", "this", "Dont", "need", "", "need"} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringExclusion{
			Name:      "MySlice",
			Blacklist: blackl,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	blackl = []string{"We", "Dont", "need", ""}
	field = []string{"This", "is", "good", "this", "Dont", "need", "", "need"} // 4 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringExclusion{
			Name:      "MySlice",
			Blacklist: blackl,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
