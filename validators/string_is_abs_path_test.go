package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsAbsPath(t *testing.T) {

	r := require.New(t)

	v := StringIsAbsPath{Name: "Name", Field: "/tmp/test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsAbsPath{Name: "Name", Field: "test"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path 'test' must be absolute"}, e.Get("Name"))

	v = StringIsAbsPath{Name: "Name", Field: "/tmp//test/test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp//test/test' must be absolute"}, e.Get("Name"))

	v = StringIsAbsPath{Name: "Name", Field: "./test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path './test' must be absolute"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsAbsPath{"Name", "test"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path 'test' must be absolute"}, e.Get("Name"))
}
