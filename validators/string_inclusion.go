package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringInclusion is a validator object
type StringInclusion struct {
	Name    string
	Field   string
	List    []string
	Message string
}

// Validate adds an error if the field is not one of the values from the list.
func (v *StringInclusion) Validate(e *validator.Errors) {

	var found = false

	for _, l := range v.List {
		if l == v.Field {
			found = true
			break
		}
	}
	if !found {
		if len(v.Message) > 0 {
			e.Add(v.Name, v.Message)
			return
		}

		e.Add(v.Name, fmt.Sprintf("%s is not in the list [%s]", v.Name, strings.Join(v.List, ", ")))
	}
}
