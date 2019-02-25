package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsLowerCase(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	v := StringIsLowerCase{Name: "Name", Field: "asdehr247"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid, spaces are trimmed
	v = StringIsLowerCase{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// outer whitespaces are allowed
	v = StringIsLowerCase{Name: "Name", Field: " wh1t3spaces "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// inner whitespaces are allowed
	v = StringIsLowerCase{Name: "Name", Field: "spac3 ins1de"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// uppercase is invalid
	v = StringIsLowerCase{Name: "Name", Field: "Abcd"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be lowercase"}, e.Get("Name"))
}
