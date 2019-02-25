package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_IntsAreEqual(t *testing.T) {

	r := require.New(t)

	v := IntsAreEqual{Name: "Number", ValueOne: 1, ValueTwo: 1}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = IntsAreEqual{Name: "Number", ValueOne: 1, ValueTwo: 2}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"1 is not equal to 2"}, e.Get("Number"))
}
