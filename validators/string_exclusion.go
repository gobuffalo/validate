package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringExclusion is a validator object.
type StringExclusion struct {
	Name      string
	Field     string
	Blacklist []string
}

// Validate adds an error if the Field is one of the values from the Blacklist.
func (v *StringExclusion) Validate(e *validator.Errors) {

	var found = false

	for _, l := range v.Blacklist {
		if l == v.Field {
			found = true
			break
		}
	}

	if found {
		e.Add(v.Name, fmt.Sprintf("%s is in the blacklist [%s]", v.Name, strings.Join(v.Blacklist, ", ")))
	}
}

// SetField sets validator field.
func (v *StringExclusion) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringExclusion) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
