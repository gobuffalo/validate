package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasLowerCase(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringHasLowerCase{Name: "Name", Field: "3w4ASF^^#"} // at least 1 lowercase
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasLowerCase{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasLowerCase{Name: "Name", Field: " space inside "} // outer and inner whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasLowerCase{Name: "Name", Field: "ABC123"} // must contain lowercase
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 lowercase"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringHasLowerCase{Name: "Name", Field: "    "} // only spaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 lowercase"}, e.Get("Name"))
}
