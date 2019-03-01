package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringMatchRegexDive(t *testing.T) {

	r := require.New(t)

	regex := Email
	field := []string{"abc@test.com", "aaa@good.too", "here@also.nice.very"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringMatchRegex{
			Name:  "MySlice",
			Regex: regex,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"abc@", "@bad.too", "h ere@also.notnice", " white@space.bad", "", " "}

	v = StringSliceDive{
		Validator: &StringMatchRegex{
			Name:  "MySlice",
			Regex: regex,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(6, e.Count())

}
