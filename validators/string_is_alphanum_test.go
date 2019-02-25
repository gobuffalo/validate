package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsAlphaNum(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsAlphaNum{Name: "Name", Field: "ASfgg5452"}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsAlphaNum{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsAlphaNum{Name: "Name", Field: "ы$^"} // any other characters except for a-zA-Z are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only numbers and/or letters"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsAlphaNum{Name: "Name", Field: " wh1t3 spaces "} // inner/outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only numbers and/or letters"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsAlphaNum{Name: "Name", Field: "   "} // only whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only numbers and/or letters"}, e.Get("Name"))
}
