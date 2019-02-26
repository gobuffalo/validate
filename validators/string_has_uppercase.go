package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringHasUpperCase is a validator object.
type StringHasUpperCase struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field contains at least 1 uppercase. Empty string is valid.
func (v *StringHasUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if rxHasUpperCase.MatchString(v.Field) {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must contain at least 1 uppercase", v.Name))
}
