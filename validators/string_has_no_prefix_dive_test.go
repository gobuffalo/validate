package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasNoPrefixDive(t *testing.T) {

	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	e := validator.NewErrorsP()
	v := StringSliceDive{
		Validator: &StringHasNoPrefix{Name: "Slice", ComparedField: "F"},
		Field:     sl,
	}
	v.Validate(e)
	r.Equal(1, e.Count()) // 4(total) - 3 strings in sl that do not have matched prefix
}
