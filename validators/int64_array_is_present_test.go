package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_Int64ArrayIsPresent(t *testing.T) {

	r := require.New(t)

	v := Int64ArrayIsPresent{Name: "Name", Field: []int64{1}}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = Int64ArrayIsPresent{Name: "Name", Field: []int64{}}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be empty"}, e.Get("Name"))

	e = validator.NewErrors()
	v = Int64ArrayIsPresent{Name: "Name", Field: []int64{}, Message: "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = Int64ArrayIsPresent{"Name", []int64{}, "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))
}
