package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_FuncValidator(t *testing.T) {

	r := require.New(t)

	fv := &FuncValidator{
		Name:  "CustomFunc",
		Field: "Field",
		Fn: func() bool {
			return false
		},
	}

	e := validator.NewErrors()
	fv.Validate(e)
	r.Equal([]string{"CustomFunc result is false"}, e.Get("CustomFunc"))
}

func Test_FuncValidatorNoName(t *testing.T) {

	r := require.New(t)

	fv := &FuncValidator{
		Field: "FuncName",
		Fn: func() bool {
			return false
		},
	}

	e := validator.NewErrors()
	fv.Validate(e)
	r.Equal([]string{"FuncName result is false"}, e.Get("FuncName"))
}
