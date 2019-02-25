package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IntsAreNotEqual(t *testing.T) {

	r := require.New(t)

	v := IntsAreNotEqual{Name: "Number", ValueOne: 2, ValueTwo: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = IntsAreNotEqual{Name: "Number", ValueOne: 2, ValueTwo: 2}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"2 is equal to 2"}, e.Get("Number"))
}
