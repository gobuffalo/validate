package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUTFLetters(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	v := StringIsUTFLetters{Name: "Name", Field: "asd品ʂля"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid
	v = StringIsUTFLetters{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// any other characters except for UTF letters are not allowed
	v = StringIsUTFLetters{Name: "Name", Field: "123~$"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode letter characters"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// inner/outer whitespaces are not allowed
	v = StringIsUTFLetters{Name: "Name", Field: " ля 品ʂ "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode letter characters"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only whitespaces are not allowed
	v = StringIsUTFLetters{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode letter characters"}, e.Get("Name"))
}
