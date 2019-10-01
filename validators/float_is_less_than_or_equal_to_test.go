package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatIsLessThanOrEqualTo(t *testing.T) {
	r := require.New(t)

	v := FloatIsLessThanOrEqualTo{Name: "Number", Field: 1, Compared: 2}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatIsLessThanOrEqualTo{Name: "Number", Field: 2, Compared: 2}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatIsLessThanOrEqualTo{Name: "number", Field: 1, Compared: 0}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"1.000000 is not less than or equal to 0.000000."})

	v = FloatIsLessThanOrEqualTo{Name: "number", Field: 1, Compared: 0, Message: "number is not less than or equal to 0."}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"number is not less than or equal to 0."})
}
