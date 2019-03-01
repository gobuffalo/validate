package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringLengthInRangeDive(t *testing.T) {

	r := require.New(t)

	min := 2
	max := 10
	field := []string{"22", "1010101010", "also", "good", "        ", "あいうえお"}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringLengthInRange{
			Name: "MySlice",
			Min:  min,
			Max:  max,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringSliceDive{
		Validator: &StringLengthInRange{
			Name: "MySlice",
			Min:  min, // only min
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringSliceDive{
		Validator: &StringLengthInRange{
			Name: "MySlice",
			Max:  max, // only max
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"1", "", "11111111111", "numb3rs", "al so"}

	v = StringSliceDive{
		Validator: &StringLengthInRange{
			Name: "MySlice",
			Min:  min,
			Max:  max,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringLengthInRange{
			Name: "MySlice",
			Min:  min, // only min
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(2, e.Count())

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringLengthInRange{
			Name: "MySlice",
			Max:  max, // only max
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())
}
