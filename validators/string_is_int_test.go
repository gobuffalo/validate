package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsInt(t *testing.T) {

	r := require.New(t)

	e := validator.NewErrors()
	v := StringIsInt{Name: "Name", Field: "1988"} // must be int
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsInt{Name: "Name", Field: "-0"} // must be int
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsInt{Name: "Name", Field: ""} // empty string is valid
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsInt{Name: "Name", Field: "13.5"} // float is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsInt{Name: "Name", Field: "baby"} // string is invalid
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsInt{Name: "Name", Field: "123 "} // whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsInt{Name: "Name", Field: "   "} // only whitespaces are not allowed
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))
}
