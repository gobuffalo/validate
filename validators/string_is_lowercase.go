package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsLowerCase is a validator object
type StringIsLowerCase struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is lowercase. Empty string is valid.
func (v *StringIsLowerCase) Validate(e *validator.Errors) {

	// null string is valid
	if IsNull(v.Field) {
		return
	}

	if v.Field == strings.ToLower(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}
	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be lowercase", v.Name))
}
