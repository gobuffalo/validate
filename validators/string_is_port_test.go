package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPort(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsPort{Name: "Name", Field: "1"} // Port is OK > 0
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPort{Name: "Name", Field: "65535"} // Port is OK < 65536
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPort{Name: "Name", Field: ""} // empty string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must represent a valid port"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsPort{Name: "Name", Field: " 13 "} // outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must represent a valid port"}, e.Get("Name"))
}
