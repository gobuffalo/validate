package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

// Possible MAC formats:
// 01:23:45:67:89:ab
// 01:23:45:67:89:ab:cd:ef
// 01-23-45-67-89-ab
// 01-23-45-67-89-ab-cd-ef
// 0123.4567.89ab
// 0123.4567.89ab.cdef
func Test_StringIsMAC(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// MAC is OK
	v := StringIsMAC{Name: "Name", Field: "01:23:45:67:89:ab"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is invalid
	v = StringIsMAC{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be valid MAC address"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// outer whitespaces are not allowed
	v = StringIsMAC{Name: "Name", Field: " 01:23:45:67:89:ab "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be valid MAC address"}, e.Get("Name"))
}
