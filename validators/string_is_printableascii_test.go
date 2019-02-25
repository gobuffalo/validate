package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPrintableASCII(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsPrintableASCII{Name: "Name", Field: "abc123"} // must be ASCII
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPrintableASCII{Name: "Name", Field: "!$#%()-=<>etc...,@"} // must be ASCII
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPrintableASCII{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPrintableASCII{Name: "Name", Field: " 123 "} // outer whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPrintableASCII{Name: "Name", Field: "   "} // only whitespaces are allowed
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPrintableASCII{Name: "Name", Field: string(rune(10))} // non-printable in invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain printable ASCII chars only"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsPrintableASCII{Name: "Name", Field: "опа"} // non-ascii in invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain printable ASCII chars only"}, e.Get("Name"))
}
