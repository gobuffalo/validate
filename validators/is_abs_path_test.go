package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IsAbsPath(t *testing.T) {

	r := require.New(t)

	v := IsAbsPath{Name: "Name", Field: "/tmp/test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = IsAbsPath{Name: "Name", Field: "test"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path 'test' must be absolute"}, e.Get("Name"))

	v = IsAbsPath{Name: "Name", Field: "/tmp//test/test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp//test/test' must be absolute"}, e.Get("Name"))

	v = IsAbsPath{Name: "Name", Field: "./test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path './test' must be absolute"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IsAbsPath{Name: "Name", Field: "test", Message: "path must be absolute"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path must be absolute"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IsAbsPath{"Name", "test", "path must be absolute"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path must be absolute"}, e.Get("Name"))
}
