package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsUpperCase is a validator object.
type StringIsUpperCase struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not uppercased. Empty string is valid.
func (v *StringIsUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if v.Field == strings.ToUpper(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be uppercase", v.Name))
}
