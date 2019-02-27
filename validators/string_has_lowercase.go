package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringHasLowerCase is a validator object
type StringHasLowerCase struct {
	Name  string
	Field string
}

// Validate adds an error if the Field contains at least 1 lowercase. Empty string is valid.
func (v *StringHasLowerCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if rxHasLowerCase.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must contain at least 1 lowercase", v.Name))
}
