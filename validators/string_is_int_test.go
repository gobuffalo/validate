package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsInt(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// must be int
	v := StringIsInt{Name: "Name", Field: "1988"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// must be int
	v = StringIsInt{Name: "Name", Field: "-0"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid
	v = StringIsInt{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// float is invalid
	v = StringIsInt{Name: "Name", Field: "13.5"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// string is invalid
	v = StringIsInt{Name: "Name", Field: "baby"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are not allowed
	v = StringIsInt{Name: "Name", Field: "123 "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only whitespaces are not allowed
	v = StringIsInt{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be an integer"}, e.Get("Name"))
}
