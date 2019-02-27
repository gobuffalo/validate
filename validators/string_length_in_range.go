package validators

import (
	"fmt"
	"unicode/utf8"

	"github.com/s3rj1k/validator"
)

// StringLengthInRange is a validator object.
type StringLengthInRange struct {
	Name  string
	Field string
	Min   int
	Max   int
}

// Validate adds an error if the Field length is not in range between Min and Max (inclusive).
// It is possible to provide either both or one of the Min/Max values.
func (v *StringLengthInRange) Validate(e *validator.Errors) {
	strLength := utf8.RuneCountInString(v.Field)
	if v.Max == 0 {
		v.Max = strLength
	}

	if strLength >= v.Min && strLength <= v.Max {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s not in range(%d, %d)", v.Name, v.Min, v.Max))
}
