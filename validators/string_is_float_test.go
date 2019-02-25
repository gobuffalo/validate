package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsFloat(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	// must be float
	v := StringIsFloat{Name: "Name", Field: "15.22"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// must be float
	v = StringIsFloat{Name: "Name", Field: "-0.0"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is valid
	v = StringIsFloat{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(0, e.Count())

	// int is valid
	v = StringIsFloat{Name: "Name", Field: "135"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// string is invalid
	v = StringIsFloat{Name: "Name", Field: "baby"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a float"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// whitespaces are not allowed
	v = StringIsFloat{Name: "Name", Field: "123.2 "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a float"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// only whitespaces are not allowed
	v = StringIsFloat{Name: "Name", Field: "   "}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a float"}, e.Get("Name"))
}
