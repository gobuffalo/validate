package validators

import (
	"fmt"
	"unicode/utf8"

	"github.com/s3rj1k/validator"
)

// StringLengthInRange is a validator object.
type StringLengthInRange struct {
	Name    string
	Field   string
	Min     int
	Max     int
	Message string
}

// Validate adds an error if the field length is not in range between min and max (inclusive).
// It is possible to provide either both or one of the min/max values.
func (v *StringLengthInRange) Validate(e *validator.Errors) {
	strLength := utf8.RuneCountInString(v.Field)
	if v.Max == 0 {
		v.Max = strLength
	}

	if v.Message == "" {
		v.Message = fmt.Sprintf("%s not in range(%d, %d)", v.Name, v.Min, v.Max)
	}

	if !(strLength >= v.Min && strLength <= v.Max) {
		e.Add(v.Name, v.Message)
	}
}
