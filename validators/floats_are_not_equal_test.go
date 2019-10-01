package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatsAreNotEqual(t *testing.T) {
	r := require.New(t)

	v := FloatsAreNotEqual{Name: "Number", ValueOne: 2, ValueTwo: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatsAreNotEqual{Name: "Number", ValueOne: 2, ValueTwo: 2}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"2.000000 is equal to 2.000000"})
}
