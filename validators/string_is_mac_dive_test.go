package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsMACDive(t *testing.T) {

	r := require.New(t)

	field := []string{"01:23:45:67:89:ab", "0a-1b-3c-4d-5e-6f", "0A:00:27:00:00:05"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsMAC{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"01:23:45:67:89:ab ", "0a-1b-3c-4d-5e-6f", "01:23:45: 67:89:ab", " ", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsMAC{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
