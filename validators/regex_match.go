package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// RegexMatch specifies the properties needed by the validation.
type RegexMatch struct {
	Name    string
	Field   string
	Expr    string
	Message string
}

// Validate performs the validation based on the regexp match.
func (v *RegexMatch) Validate(e *validator.Errors) {
	r := regexp.MustCompile(v.Expr)
	if r.Match([]byte(v.Field)) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s does not match the expected format.", v.Name))
}
