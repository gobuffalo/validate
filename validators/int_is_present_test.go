package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IntIsPresent(t *testing.T) {

	r := require.New(t)

	v := IntIsPresent{Name: "Name", Field: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = IntIsPresent{Name: "Name", Field: 0}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IntIsPresent{Name: "Name"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be blank"}, e.Get("Name"))
}
