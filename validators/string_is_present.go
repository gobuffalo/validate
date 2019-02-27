package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsPresent is a validator object.
type StringIsPresent struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is empty or has only whitespaces.
// If you don't want whitespaces - see StringIsNull validator.
func (v *StringIsPresent) Validate(e *validator.Errors) {
	if strings.TrimSpace(v.Field) != "" {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank", v.Name))
}
