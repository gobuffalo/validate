package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringsAreNotEqual is a validator object.
type StringsAreNotEqual struct {
	Name     string
	Field    string
	Compared string
	Message  string
}

// Validate adds an error if the field and compared are equal.
func (v *StringsAreNotEqual) Validate(e *validator.Errors) {
	if strings.TrimSpace(v.Field) == strings.TrimSpace(v.Compared) {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s is equal to %s", v.Field, v.Compared)
		}

		e.Add(v.Name, v.Message)
	}
}
