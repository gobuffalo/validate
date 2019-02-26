package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringsAreEqual is a validator object.
type StringsAreEqual struct {
	Name            string
	Field           string
	ComparedName    string
	ComparedField   string
	CaseInsensitive bool
	Message         string
}

// Validate adds an error if the field and compared are not equal.
func (v *StringsAreEqual) Validate(e *validator.Errors) {

	var caseName string

	if v.CaseInsensitive {
		caseName = "iequal"
		if strings.EqualFold(v.Field, v.ComparedField) {
			return
		}
	} else {
		caseName = "equal"
		if v.Field == v.ComparedField {
			return
		}
	}

	if v.Message != "" {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s does not %s %s", v.Field, caseName, v.ComparedField))
}

// SetField sets validator field.
func (v *StringsAreEqual) SetField(s string) {
	v.Field = s
}
