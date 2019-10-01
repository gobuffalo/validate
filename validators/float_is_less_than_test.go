package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatIsLessThan(t *testing.T) {
	r := require.New(t)

	v := FloatIsLessThan{Name: "Number", Field: 1, Compared: 2}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatIsLessThan{Name: "number", Field: 1, Compared: 0}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"1.000000 is not less than 0.000000."})

	v = FloatIsLessThan{Name: "number", Field: 1, Compared: 0, Message: "number is not less than 0."}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"number is not less than 0."})
}
