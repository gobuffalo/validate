package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_UintArrayIsPresent(t *testing.T) {

	r := require.New(t)

	v := UintArrayIsPresent{Name: "Name", Field: []uint{1}}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = UintArrayIsPresent{Name: "Name", Field: []uint{}}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be empty"}, e.Get("Name"))

	e = validator.NewErrors()
	v = UintArrayIsPresent{"Name", []uint{}}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be empty"}, e.Get("Name"))
}
