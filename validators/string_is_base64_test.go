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
	v := StringIsBase64{Name: "Name", Field: sEnc} // must be base64 string
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsBase64{Name: "Name", Field: ""} // empty string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsBase64{Name: "Name", Field: " " + sEnc + " "} // outer whitespaces are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsBase64{Name: "Name", Field: "   "} // only whitespaces are invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	e = validator.NewErrors()
	sEnc32 := base32.StdEncoding.EncodeToString([]byte("abc123"))
	v = StringIsBase64{Name: "Name", Field: sEnc32} // base32 is bad
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsBase64{Name: "Name", Field: "abc123"} // simple string is bad
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))
}
