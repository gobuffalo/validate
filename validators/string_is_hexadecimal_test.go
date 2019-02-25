package validators

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsHexadecimal(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsHexadecimal{Name: "Name", Field: hex.EncodeToString([]byte("Hello"))} // hex here
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsHexadecimal{Name: "Name", Field: "FF"} // hex here also
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsHexadecimal{Name: "Name", Field: fmt.Sprintf("%x", 155)} // and this is hex
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsHexadecimal{Name: "Name", Field: "Hello"} // other strings are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsHexadecimal{Name: "Name", Field: ""} // empty string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsHexadecimal{Name: "Name", Field: "    "} // whitespaces are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsHexadecimal{Name: "Name", Field: "FF "} // whitespaces are not trimmed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))
}
