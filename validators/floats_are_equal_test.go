package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatsAreEqual(t *testing.T) {
	r := require.New(t)

	v := FloatsAreEqual{Name: "Number", ValueOne: 1, ValueTwo: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatsAreEqual{Name: "Number", ValueOne: 1, ValueTwo: 2}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"1.000000 is not equal to 2.000000"})
}
