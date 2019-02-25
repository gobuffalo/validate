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

	v = StringIsLowerCase{Name: "Name", Field: "   "} // empty string is valid, spaces are trimmed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsLowerCase{Name: "Name", Field: " wh1t3spaces "} // outer whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsLowerCase{Name: "Name", Field: "spac3 ins1de"} // inner whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsLowerCase{Name: "Name", Field: "Abcd"} // uppercase is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be lowercase"}, e.Get("Name"))
}
