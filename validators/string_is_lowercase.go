package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsLowerCase is a validator object.
type StringIsLowerCase struct {
	Name  string
	Field string
}

// Validate adds an error if the field is not lowercased. Empty string is valid.
func (v *StringIsLowerCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if v.Field == strings.ToLower(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be lowercase", v.Name))
}
