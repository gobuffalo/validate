package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IsDir(t *testing.T) {

	r := require.New(t)

	err := os.MkdirAll("/tmp/test_dir", 0666)
	r.Nil(err)

	v := IsDir{Name: "Name", Field: "/tmp/test_dir"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove("/tmp/test_dir")
	r.Nil(err)

	v = IsDir{Name: "Name", Field: "/tmp/test_dir"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp/test_dir' is not dir"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IsDir{Name: "Name", Field: "/tmp/test_dir", Message: "path is not dir"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path is not dir"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IsDir{"Name", "/tmp/test_dir", "path is not dir"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path is not dir"}, e.Get("Name"))
}
