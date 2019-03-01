package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringInclusionDive(t *testing.T) {

	r := require.New(t)

	whitel := []string{"This", "is", "good", "this", "Dont", "ne ed", "", "need"}
	field := []string{"This", "is", "good", "this", "Dont", "ne ed", "", "need"} // 0 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringInclusion{
			Name:      "MySlice",
			Whitelist: whitel,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	whitel = []string{"Only", "These"}
	field = []string{"These", "Only", "good", "Only", "These", "", "need"} // 3 errors

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringInclusion{
			Name:      "MySlice",
			Whitelist: whitel,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())
}
