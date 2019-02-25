package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsASCII(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// must be ASCII
	v := StringIsASCII{Name: "Name", Field: "abc123"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// must be ASCII
	v = StringIsASCII{Name: "Name", Field: "!$#%()-=<>etc...,@"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid
	v = StringIsASCII{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// outer whitespaces are allowed
	v = StringIsASCII{Name: "Name", Field: " 123 "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// only whitespaces are allowed
	v = StringIsASCII{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// non-ascii in invalid
	v = StringIsASCII{Name: "Name", Field: "опа"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain ASCII chars only"}, e.Get("Name"))
}
