package validators

import (
	"fmt"
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
