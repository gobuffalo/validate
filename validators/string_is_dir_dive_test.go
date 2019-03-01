package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsDirDive(t *testing.T) {

	r := require.New(t)

	err := os.MkdirAll("/tmp/test_dir", 0666)
	r.Nil(err)

	err = os.MkdirAll("/tmp/test_dir2", 0666)
	r.Nil(err)

	field := []string{"/tmp/test_dir", "/tmp/test_dir2"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsDir{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove("/tmp/test_dir")
	r.Nil(err)

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsDir{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.Remove("/tmp/test_dir2")
	r.Nil(err)

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsDir{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(2, e.Count())
}
