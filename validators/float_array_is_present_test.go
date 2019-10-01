package validators_test

import (
	"testing"

	"github.com/gobuffalo/validate"
	. "github.com/gobuffalo/validate/validators"
	"github.com/stretchr/testify/require"
)

func Test_FloatArrayIsPresent(t *testing.T) {
	r := require.New(t)

	v := FloatArrayIsPresent{Name: "Name", Field: []float64{1}}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = FloatArrayIsPresent{Name: "Name", Field: []float64{}}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Name can not be empty."})

	errors = validate.NewErrors()
	v = FloatArrayIsPresent{Name: "Name", Field: []float64{}, Message: "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})

	errors = validate.NewErrors()
	v = FloatArrayIsPresent{"Name", []float64{}, "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})
}
