package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_BytesArePresent(t *testing.T) {

	r := require.New(t)

	v := BytesArePresent{Name: "Name", Field: []byte("Mark")}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = BytesArePresent{Name: "Name", Field: []byte(" ")}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = BytesArePresent{Name: "Name", Field: []byte("")}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be blank"}, e.Get("Name"))
}
