package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIP(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsIP{Name: "Name", Field: "127.0.0.1"} // IPV4 is OK
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsIP{Name: "Name", Field: "2001:db8:6:56::53"} // IPV6 is OK
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsIP{Name: "Name", Field: ""} // empty string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be either IP version 4 or 6"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsIP{Name: "Name", Field: "12a.asd.12"} // random string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be either IP version 4 or 6"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsIP{Name: "Name", Field: " 127.0.0.1 "} // outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be either IP version 4 or 6"}, e.Get("Name"))
}
