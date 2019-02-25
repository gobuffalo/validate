package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsASCII(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsASCII{Name: "Name", Field: "abc123"} // must be ASCII
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsASCII{Name: "Name", Field: "!$#%()-=<>etc...,@"} // must be ASCII
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsASCII{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsASCII{Name: "Name", Field: " 123 "} // outer whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsASCII{Name: "Name", Field: "   "} // only whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsASCII{Name: "Name", Field: "опа"} // non-ascii in invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain ASCII chars only"}, e.Get("Name"))
}
