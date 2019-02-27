package validators

import (
	"fmt"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFLetterNum is a validator object.
type StringIsUTFLetterNum struct {
	Name  string
	Field string
}

// Validate adds an error if the field contains anything except unicode letters/numbers (category L/N).
// Empty string is valid.
func (v *StringIsUTFLetterNum) Validate(e *validator.Errors) {
	var badRune bool

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// checking each rune
	for _, c := range v.Field {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			badRune = true
			break
		}
	}

	if badRune {
		e.Add(v.Name, fmt.Sprintf("%s must contain only unicode letter/number characters", v.Name))
	}
}
