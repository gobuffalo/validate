package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IsPath(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/test") // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	defer func(r *require.Assertions) {
		err = os.Remove("/tmp/test")
		r.Nil(err)
	}(r)

	v := IsPath{Name: "Name", Path: "/tmp/test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = IsPath{Name: "Name", Path: "/tmp/doesnotexist"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp/doesnotexist' must exist"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IsPath{Name: "Name", Path: "", Message: "path must exist"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path must exist"}, e.Get("Name"))

	e = validator.NewErrors()
	v = IsPath{"Name", "", "path must exist"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path must exist"}, e.Get("Name"))
}
