package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPDive(t *testing.T) {

	r := require.New(t)

	field := []string{"684D:1111:222:3333:4444:5555:6:77", "0:0:0:0:0:0:0:0", "2001:db8:1::ab9:C0A8:102", "255.255.253.0", "19.117.63.126"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsIP{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"265.255.253.0", "127.0.0.", "0:0:0:0:0:0", " 127.0.0.1", "127.0. 0.1", " ", ""}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsIP{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(7, e.Count())
}
