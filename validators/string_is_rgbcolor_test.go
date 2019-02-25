package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsRGBcolor(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsRGBcolor{Name: "Name", Field: "rgb(0,0,0)"} // hexcolor here
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsRGBcolor{Name: "Name", Field: "rgb(255,255,255)"} // hexcolor here also (3-6 0-F chars)
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsRGBcolor{Name: "Name", Field: "RGB(0,15,25)"} // rgb must be lowercased
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsRGBcolor{Name: "Name", Field: "rgb(0,0,256)"} // values 0-255
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsRGBcolor{Name: "Name", Field: "    "} // empty string or only whitespaces are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsRGBcolor{Name: "Name", Field: "ffd762 "} // whitespaces are not trimmed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))
}
