package validators

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFNumeric is a validator object
type StringIsUTFNumeric struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string contains only unicode numbers of any kind (category N).
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩. Empty string is valid.
func (v *StringIsUTFNumeric) Validate(e *validator.Errors) {

	var failed bool
	var field = v.Field

	// null string is valid
	if IsNull(field) {
		return
	}

	if strings.IndexAny(field, "+-") > 0 {
		failed = true
	}

	if len(field) > 1 {
		field = strings.TrimPrefix(field, "-")
		field = strings.TrimPrefix(field, "+")
	}

	for _, c := range field {
		if !unicode.IsNumber(c) { //numbers && minus sign are ok
			failed = true
		}
	}

	if !failed {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain only unicode numbers", v.Name))
}
