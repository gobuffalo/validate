package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringMatchRegex is a validator object.
type StringMatchRegex struct {
	Name    string
	Field   string
	Regex   string
	Message string
}

// Validate adds an error if the field does not match regular expression expr.
func (v *StringMatchRegex) Validate(e *validator.Errors) {
	r := regexp.MustCompile(v.Regex)
	if r.Match([]byte(v.Field)) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s does not match the expected format", v.Name))
}
