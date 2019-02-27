package validators

import (
	"fmt"
	"regexp"
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
}

// Validate adds an error if the Field is not equal to ComparedField.
// CaseInsensitive flag can be set to make comparison case insensitive.
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

	if len(v.ComparedName) == 0 {
		e.Add(v.Name, fmt.Sprintf("%s does not %s %s", v.Field, caseName, v.ComparedField))
	} else {
		e.Add(v.Name, fmt.Sprintf("%s does not %s %s", v.Name, caseName, v.ComparedName))
	}
}

// SetField sets validator field.
func (v *StringsAreEqual) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringsAreEqual) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
