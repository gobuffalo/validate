package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsJSON(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	v := StringIsJSON{Name: "Name", Field: "{\"test\": \"sure\"}"}
	v.Validate(e)
	r.Equal(0, e.Count())

	// outer whitespaces are valid
	v = StringIsJSON{Name: "Name", Field: "   {\"test\": \"sure\"}    "}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty string is invalid
	v = StringIsJSON{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be a valid JSON"}, e.Get("Name"))
}
