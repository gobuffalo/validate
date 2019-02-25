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

	// empty string is valid
	v = StringIsAlpha{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// any other characters except for a-zA-Z are invalid
	v = StringIsAlpha{Name: "Name", Field: " 1$~复制品"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only letters (a-zA-Z)"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// inner/outer whitespaces are not allowed
	v = StringIsAlpha{Name: "Name", Field: " white spaces "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only letters (a-zA-Z)"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only whitespaces are not allowed
	v = StringIsAlpha{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only letters (a-zA-Z)"}, e.Get("Name"))
}
