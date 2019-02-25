package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsRGBcolor(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// hexcolor here
	v := StringIsRGBcolor{Name: "Name", Field: "rgb(0,0,0)"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// hexcolor here also (3-6 0-F chars)
	v = StringIsRGBcolor{Name: "Name", Field: "rgb(255,255,255)"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// rgb must be lowercased
	v = StringIsRGBcolor{Name: "Name", Field: "RGB(0,15,25)"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// values 0-255
	v = StringIsRGBcolor{Name: "Name", Field: "rgb(0,0,256)"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// empty string or only whitespaces are invalid
	v = StringIsRGBcolor{Name: "Name", Field: "    "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are not trimmed
	v = StringIsRGBcolor{Name: "Name", Field: "ffd762 "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a RGB color in form rgb(RRR, GGG, BBB)"}, e.Get("Name"))
}
