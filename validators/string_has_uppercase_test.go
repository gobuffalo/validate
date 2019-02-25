package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasUpperCase(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// at least 1 upper case
	v := StringHasUpperCase{Name: "Name", Field: "3w4asF@`"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid
	v = StringHasUpperCase{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// outer and inner whitespaces are allowed
	v = StringHasUpperCase{Name: "Name", Field: " Space Inside "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// must contain uppercase
	v = StringHasUpperCase{Name: "Name", Field: "abc123"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 uppercase"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only spaces are not allowed
	v = StringHasUpperCase{Name: "Name", Field: "    "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 uppercase"}, e.Get("Name"))
}
