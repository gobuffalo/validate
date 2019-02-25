package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsMAC(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsMAC{Name: "Name", Field: "01:23:45:67:89:ab"} // MAC is OK
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsMAC{Name: "Name", Field: ""} // empty string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be valid MAC address"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsMAC{Name: "Name", Field: " 01:23:45:67:89:ab "} // outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be valid MAC address"}, e.Get("Name"))
}
