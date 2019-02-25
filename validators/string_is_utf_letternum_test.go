package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUTFLetterNum(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsUTFLetterNum{Name: "Name", Field: "a১522௫sd品ʂля٣"}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUTFLetterNum{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUTFLetterNum{Name: "Name", Field: ":~$"} // any other characters except for UTF letters are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode letter/number characters"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsUTFLetterNum{Name: "Name", Field: " ля 品ʂ "} // inner/outer whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode letter/number characters"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsUTFLetterNum{Name: "Name", Field: "   "} // only whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain only unicode letter/number characters"}, e.Get("Name"))
}
