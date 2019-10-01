package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatIsPresent(t *testing.T) {
	r := require.New(t)

	v := FloatIsPresent{Name: "Name", Field: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatIsPresent{Name: "Name", Field: 0}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Name can not be blank."})

	errors = validate.NewErrors()
	v = FloatIsPresent{Name: "Name", Field: 0, Message: "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})

	errors = validate.NewErrors()
	v = FloatIsPresent{"Name", 0, "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})
}
