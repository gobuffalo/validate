package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsURLDive(t *testing.T) {

	r := require.New(t)

	field := []string{"", "http://", "https://", "http", "google.com",
		"http://www.google.com", "http://google.com", "https://www.google.cOM",
		"ht123tps://www.google.cOM",
		"https://golang.Org",
		"https://invalid#$%#$@.Org"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsURL{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(7, e.Count())
}
