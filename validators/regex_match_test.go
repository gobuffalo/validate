package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_RegexMatch(t *testing.T) {

	r := require.New(t)

	v := RegexMatch{Name: "Phone", Field: "555-555-5555", Expr: "^([0-9]{3}-[0-9]{3}-[0-9]{4})$"}
	e := validator.NewErrors()
	v.Validate(e)

	r.Equal(0, e.Count())

	v = RegexMatch{Name: "Phone", Field: "123-ab1-1424", Expr: "^([0-9]{3}-[0-9]{3}-[0-9]{4})$"}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Phone does not match the expected format"}, e.Get("Phone"))

	e = validator.NewErrors()
	v = RegexMatch{Name: "Phone", Field: "123-ab1-1424", Expr: "^([0-9]{3}-[0-9]{3}-[0-9]{4})$", Message: "Phone number does not match the expected format"}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Phone number does not match the expected format"}, e.Get("Phone"))

	e = validator.NewErrors()
	v = RegexMatch{"Phone", "123-ab1-1424", "^([0-9]{3}-[0-9]{3}-[0-9]{4})$", "Phone number does not match the expected format"}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Phone number does not match the expected format"}, e.Get("Phone"))
}
