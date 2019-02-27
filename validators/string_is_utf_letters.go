package validators

import (
	"fmt"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFLetters is a validator object.
type StringIsUTFLetters struct {
	Name  string
	Field string
}

// Validate adds an error if the field contains anything except unicode letters (category L)
// Similar to StringIsAlpha but for all languages. Empty string is valid.
func (v *StringIsUTFLetters) Validate(e *validator.Errors) {
	var badRune bool

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// checking each rune
	for _, c := range v.Field {
		if !unicode.IsLetter(c) {
			badRune = true
			break
		}
	}

	if badRune {
		e.Add(v.Name, fmt.Sprintf("%s must contain only unicode letter characters", v.Name))
	}
}
