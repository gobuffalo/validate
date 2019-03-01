package validators

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFLetters is a validator object.
type StringIsUTFLetters struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains anything except unicode letters (category L)
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

// SetField sets validator field.
func (v *StringIsUTFLetters) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUTFLetters) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
