package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringsMatch is a validator object
type StringsMatch struct {
	Name    string
	Field   string
	Field2  string
	Message string
}

// Validate performs the validation equality of two strings.
func (v *StringsMatch) Validate(e *validator.Errors) {
	if strings.TrimSpace(v.Field) != strings.TrimSpace(v.Field2) {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not equal %s", v.Field, v.Field2)
		}

		e.Add(v.Name, v.Message)
	}
}
