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

	// hex here
	v := StringIsHexadecimal{Name: "Name", Field: hex.EncodeToString([]byte("Hello"))}
	v.Validate(e)
	r.Equal(0, e.Count())

	// hex here also
	v = StringIsHexadecimal{Name: "Name", Field: "FF"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// and this is hex
	v = StringIsHexadecimal{Name: "Name", Field: fmt.Sprintf("%x", 155)}
	v.Validate(e)
	r.Equal(0, e.Count())

	// other strings are invalid
	v = StringIsHexadecimal{Name: "Name", Field: "Hello"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// empty string is invalid
	v = StringIsHexadecimal{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are invalid
	v = StringIsHexadecimal{Name: "Name", Field: "    "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are not trimmed
	v = StringIsHexadecimal{Name: "Name", Field: "FF "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a hexadecimal number"}, e.Get("Name"))
}
