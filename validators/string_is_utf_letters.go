package validators

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/s3rj1k/validator"
)

// StringIsUTFLetters is a validator object
type StringIsUTFLetters struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string contains only unicode letter characters (category L).
// Similar to StringIsAlpha but for all languages. Empty string is valid.
func (v *StringIsUTFLetters) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// checking each rune
	for i, c := range v.Field {
		if !unicode.IsLetter(c) {
			break
		}
		if i == utf8.RuneCountInString(v.Field)-1 {
			return
		}
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain only unicode letter characters", v.Name))
}
