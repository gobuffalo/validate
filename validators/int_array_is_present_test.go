package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IntArrayIsPresent(t *testing.T) {

	r := require.New(t)

	v := IntArrayIsPresent{Name: "Name", Field: []int{1}}
	e := validator.NewErrors()
	v.Validate(e)

	r.Equal(0, e.Count())

	v = IntArrayIsPresent{Name: "Name", Field: []int{}}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be empty."}, e.Get("Name"))

	e = validator.NewErrors()
	v = IntArrayIsPresent{Name: "Name", Field: []int{}, Message: "Field can't be blank."}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank."}, e.Get("Name"))

	e = validator.NewErrors()
	v = IntArrayIsPresent{"Name", []int{}, "Field can't be blank."}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank."}, e.Get("Name"))
}
