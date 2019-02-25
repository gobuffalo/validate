package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPresent(t *testing.T) {

	r := require.New(t)

	v := StringIsPresent{Name: "Name", Field: "Mark"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPresent{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsPresent{Name: "Name", Field: "", Message: "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsPresent{"Name", "", "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))
}
