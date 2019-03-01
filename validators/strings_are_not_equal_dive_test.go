package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringAreNotEqualDive(t *testing.T) {

	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	e := validator.NewErrorsP()
	v := StringSliceDive{
		Validator: &StringsAreEqual{Name: "Slice", ComparedField: "Bar"},
		Field:     sl,
	}
	v.Validate(e)
	r.Equal(3, e.Count())

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringsAreEqual{Name: "Slice", ComparedField: "Bar"},
		Field:     []string{"Bar", "Bar", "bar", ""},
	}
	v.Validate(e)
	r.Equal(2, e.Count())
}
