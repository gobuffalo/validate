package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsAbsPathDive(t *testing.T) {

	r := require.New(t)

	field := []string{"/var/log", "/usr/local/bin", "/tmp/test", "/tmp//test/test", "./test", "test", ""} // 4 errors

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsAbsPath{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
