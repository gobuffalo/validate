package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasUpperCase(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringHasUpperCase{Name: "Name", Field: "3w4asF@`"} // at least 1 upper case
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasUpperCase{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasUpperCase{Name: "Name", Field: " Space Inside "} // outer and inner whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasUpperCase{Name: "Name", Field: "abc123"} // must contain uppercase
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 uppercase"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringHasUpperCase{Name: "Name", Field: "    "} // only spaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 uppercase"}, e.Get("Name"))
}
