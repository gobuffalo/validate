package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNull(t *testing.T) {

	r := require.New(t)

	v := StringIsNull{Name: "Name", Field: ""}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsNull{Name: "Name", Field: *new(string)}
	v.Validate(e)
	r.Equal(0, e.Count())

	e = validator.NewErrors()
	v = StringIsNull{Name: "Name", Field: " ", Message: "Field must be empty"} // whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field must be empty"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsNull{"Name", ",", "Field must be empty"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field must be empty"}, e.Get("Name"))
}
