package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringInclusion(t *testing.T) {

	r := require.New(t)

	l := []string{"Mark", "Bates"}

	e := validator.NewErrors()
	v := StringInclusion{Name: "Name", Field: "Mark", List: l}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringInclusion{Name: "Name", Field: "Foo", List: l}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name is not in the list [Mark, Bates]"}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringInclusion{"Name", "Foo", l}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name is not in the list [Mark, Bates]"}, e.Get("Name"))
}
