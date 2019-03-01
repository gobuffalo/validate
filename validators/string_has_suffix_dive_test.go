package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasSuffixDive(t *testing.T) {

	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	e := validator.NewErrorsP()
	v := StringSliceDive{
		Validator: &StringHasSuffix{Name: "Slice", ComparedField: "r"},
		Field:     sl,
	}
	v.Validate(e)
	r.Equal(1, e.Count()) // 1 strings in sl that do not have matched suffix
}
