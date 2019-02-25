package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_Int64IsPresent(t *testing.T) {

	r := require.New(t)

	v := Int64IsPresent{Name: "Name", Field: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = Int64IsPresent{Name: "Name", Field: 0}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = Int64IsPresent{Name: "Name", Field: 0, Message: "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = Int64IsPresent{"Name", 0, "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))
}
