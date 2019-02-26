package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringsAreEqual is a validator object.
type StringsAreEqual struct {
	Name     string
	Field    string
	Compared string
	Message  string
}

// Validate adds an error if the field and compared are not equal.
func (v *StringsAreEqual) Validate(e *validator.Errors) {
	if strings.TrimSpace(v.Field) != strings.TrimSpace(v.Compared) {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not equal to %s", v.Field, v.Compared)
		}

		e.Add(v.Name, v.Message)
	}
}
