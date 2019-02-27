package validators

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFNumeric is a validator object
type StringIsUTFNumeric struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains anything except unicode numbers (category N).
// Leading sign is allowed. Empty string is valid.
func (v *StringIsUTFNumeric) Validate(e *validator.Errors) {

	var failed bool
	var field = v.Field

	// null string is valid
	if isNullString(field) {
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

	if failed {
		e.Add(v.Name, fmt.Sprintf("%s must contain only unicode numbers", v.Name))
	}
}
