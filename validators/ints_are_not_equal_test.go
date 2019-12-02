package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate/v2"
	. "github.com/gobuffalo/validate/v2/validators"
	"github.com/stretchr/testify/require"
)

func Test_IntsAreNotEqual(t *testing.T) {
	r := require.New(t)

	v := IntsAreNotEqual{Name: "Number", ValueOne: 2, ValueTwo: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = IntsAreNotEqual{Name: "Number", ValueOne: 2, ValueTwo: 2}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"2 is equal to 2"})
}
