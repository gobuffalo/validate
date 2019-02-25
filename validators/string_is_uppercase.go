package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsUpperCase is a validator object
type StringIsUpperCase struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is uppercase. Empty string is valid.
func (v *StringIsUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if v.Field == strings.ToUpper(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be uppercase", v.Name))
}
