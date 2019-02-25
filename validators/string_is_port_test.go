package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPort(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// Port is OK > 0
	v := StringIsPort{Name: "Name", Field: "1"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// Port is OK < 65536
	v = StringIsPort{Name: "Name", Field: "65535"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is invalid
	v = StringIsPort{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must represent a valid port"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// outer whitespaces are not allowed
	v = StringIsPort{Name: "Name", Field: " 13 "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must represent a valid port"}, e.Get("Name"))
}
