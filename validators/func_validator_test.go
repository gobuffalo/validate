package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_FuncValidator(t *testing.T) {

	r := require.New(t)

	fv := &FuncValidator{
		Name:    "Name",
		Field:   "Field",
		Message: "%s is an invalid name",
		Fn: func() bool {
			return false
		},
	}

	e := validator.NewErrors()
	fv.Validate(e)

	r.Equal([]string{"Field is an invalid name"}, e.Get("Name"))
}

func Test_FuncValidatorNoName(t *testing.T) {

	r := require.New(t)

	fv := &FuncValidator{
		Field:   "Name",
		Message: "%s is invalid",
		Fn: func() bool {
			return false
		},
	}

	e := validator.NewErrors()
	fv.Validate(e)

	r.Equal([]string{"Name is invalid"}, e.Get("Name"))
}
