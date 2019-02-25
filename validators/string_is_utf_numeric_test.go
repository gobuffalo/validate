package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUTFNumeric(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsUTFNumeric{Name: "Name", Field: "১522௫٣"}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUTFNumeric{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUTFNumeric{Name: "Name", Field: "ag:~$"} // any other characters except for UTF letters are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode numbers"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsUTFNumeric{Name: "Name", Field: " ля 品ʂ "} // inner/outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode numbers"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsUTFNumeric{Name: "Name", Field: "   "} // only whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode numbers"}, e.Get("Name"))
}
