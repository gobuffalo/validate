package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringExclusion is a validator object.
type StringExclusion struct {
	Name      string
	Field     string
	Blacklist []string
}

// Validate adds an error if the field is one of the values from the blacklist.
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
