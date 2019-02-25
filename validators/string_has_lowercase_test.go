package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasLowerCase(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// at least 1 lowercase
	v := StringHasLowerCase{Name: "Name", Field: "3w4ASF^^#"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid
	v = StringHasLowerCase{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// outer and inner whitespaces are allowed
	v = StringHasLowerCase{Name: "Name", Field: " space inside "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// must contain lowercase
	v = StringHasLowerCase{Name: "Name", Field: "ABC123"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 lowercase"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only spaces are not allowed
	v = StringHasLowerCase{Name: "Name", Field: "    "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 lowercase"}, e.Get("Name"))
}
