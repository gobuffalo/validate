package validators

import (
	"encoding/base32"
	"encoding/base64"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsBase64(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	sEnc := base64.StdEncoding.EncodeToString([]byte("abc123"))

	// must be base64 string
	v := StringIsBase64{Name: "Name", Field: sEnc}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is invalid
	v = StringIsBase64{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// outer whitespaces are invalid
	v = StringIsBase64{Name: "Name", Field: " " + sEnc + " "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only whitespaces are invalid
	v = StringIsBase64{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	sEnc32 := base32.StdEncoding.EncodeToString([]byte("abc123"))

	// base32 is bad
	v = StringIsBase64{Name: "Name", Field: sEnc32}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// simple string is bad
	v = StringIsBase64{Name: "Name", Field: "abc123"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))
}
