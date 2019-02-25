package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUpperCase(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsUpperCase{Name: "Name", Field: "ASFADG44"}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: "   "} // empty string is valid, spaces are trimmed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: " A5555 "} // outer whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: "AD GGGG"} // inner whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: "Abcd"} // lowercase is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be uppercase"}, e.Get("Name"))
}
