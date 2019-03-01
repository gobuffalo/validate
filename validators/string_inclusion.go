package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringInclusion is a validator object
type StringInclusion struct {
	Name      string
	Field     string
	Whitelist []string
}

// Validate adds an error if the Field is NOT one of the values from the Whitelist.
func (v *StringInclusion) Validate(e *validator.Errors) {

	var found = false

	for _, l := range v.Whitelist {
		if l == v.Field {
			found = true
			break
		}
	}

	if !found {
		e.Add(v.Name, fmt.Sprintf("%s is not in the list [%s]", v.Name, strings.Join(v.Whitelist, ", ")))
	}
}

// SetField sets validator field.
func (v *StringInclusion) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringInclusion) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
