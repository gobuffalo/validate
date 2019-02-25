package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsAlpha(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsAlpha{Name: "Name", Field: "abcZXC"}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsAlpha{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsAlpha{Name: "Name", Field: " 1$~复制品"} // any other characters except for a-zA-Z are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only letters (a-zA-Z)"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsAlpha{Name: "Name", Field: " white spaces "} // inner/outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only letters (a-zA-Z)"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsAlpha{Name: "Name", Field: "   "} // only whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only letters (a-zA-Z)"}, e.Get("Name"))
}
