package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringHasUpperCase is a validator object.
type StringHasUpperCase struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains at least 1 uppercase. Empty string is valid.
func (v *StringHasUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if rxHasUpperCase.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must contain at least 1 uppercase", v.Name))
}
