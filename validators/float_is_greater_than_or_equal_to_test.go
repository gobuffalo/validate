package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatIsGreaterThanOrEqualTo(t *testing.T) {
	r := require.New(t)

	v := FloatIsGreaterThanOrEqualTo{Name: "Number", Field: 2, Compared: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(0, errors.Count())

	v = FloatIsGreaterThanOrEqualTo{Name: "Number", Field: 2, Compared: 2}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(0, errors.Count())

	v = FloatIsGreaterThanOrEqualTo{Name: "number", Field: 1, Compared: 2}
	v.IsValid(errors)
	r.Equal(1, errors.Count())
	r.Equal(errors.Get("number"), []string{"1.000000 is not greater than or equal to 2.000000."})

	v = FloatIsGreaterThanOrEqualTo{Name: "number", Field: 1, Compared: 2, Message: "number isn't greater than or equal to 2."}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(1, errors.Count())
	r.Equal(errors.Get("number"), []string{"number isn't greater than or equal to 2."})
}
