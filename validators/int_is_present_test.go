package validators

import (
	"testing"

	"github.com/gobuffalo/validate/v3"
	"github.com/stretchr/testify/require"
)

func Test_IntIsPresent(t *testing.T) {
	r := require.New(t)

	v := IntIsPresent{Name: "Name", Field: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = IntIsPresent{Name: "Name", Field: 0}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Name can not be blank."})

	errors = validate.NewErrors()
	v = IntIsPresent{Name: "Name", Field: 0, Message: "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})
}
