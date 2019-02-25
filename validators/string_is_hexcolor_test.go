package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsHexcolor(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// hexcolor here
	v := StringIsHexcolor{Name: "Name", Field: "#b8f2b2"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// hexcolor here also (3-6 0-F chars)
	v = StringIsHexcolor{Name: "Name", Field: "#f00"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// must start with #
	v = StringIsHexcolor{Name: "Name", Field: "f00"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal color"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// empty string is invalid
	v = StringIsHexcolor{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal color"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are invalid
	v = StringIsHexcolor{Name: "Name", Field: "    "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal color"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are not trimmed
	v = StringIsHexcolor{Name: "Name", Field: "ffd762 "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal color"}, e.Get("Name"))
}

/*
Hex is valid : #1f1f1F , true
Hex is valid : #AFAFAF , true
Hex is valid : #1AFFa1 , true
Hex is valid : #222fff , true
Hex is valid : #F00 , true
Hex is valid : 123456 , false
Hex is valid : #afafah , false
Hex is valid : #123abce , false
Hex is valid : aFaE3f , false
Hex is valid : F00 , false
Hex is valid : #afaf , false
Hex is valid : #F0h , false
*/
