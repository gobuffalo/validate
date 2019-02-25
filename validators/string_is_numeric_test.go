package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNumeric(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsNumeric{Name: "Name", Field: "0123456789"}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsNumeric{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsNumeric{Name: "Name", Field: "abc123"} // any other characters except for a-zA-Z are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only numbers"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsNumeric{Name: "Name", Field: " 123 123 "} // inner/outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only numbers"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsNumeric{Name: "Name", Field: "  "} // only whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only numbers"}, e.Get("Name"))
}
