package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

var rxEmail *regexp.Regexp // nolint: gochecknoglobals

func init() { // nolint: gochecknoinits
	rxEmail = regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$") // nolint: lll,megacheck
}

// EmailIsPresent is a validator object
type EmailIsPresent struct {
	Name    string
	Field   string
	Message string
}

// Validate performs the validation based on the email regexp match.
func (v *EmailIsPresent) Validate(e *validator.Errors) {
	if !rxEmail.Match([]byte(v.Field)) {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not match the email format.", v.Name)
		}

		e.Add(v.Name, v.Message)
	}
}

// EmailLike checks that email has two parts (username and domain separated by @)
// Also it check that domain have domain zone (don`t check that zone is valid)
type EmailLike struct {
	Name    string
	Field   string
	Message string
}

// Validate performs the validation based on email struct (username@domain)
func (v *EmailLike) Validate(e *validator.Errors) {
	parts := strings.Split(v.Field, "@")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not match the email format.", v.Name)
		}

		e.Add(v.Name, v.Message)
	} else if len(parts) == 2 {
		domain := parts[1]
		// Check that domain is valid
		if len(strings.Split(domain, ".")) < 2 {
			if v.Message == "" {
				v.Message = fmt.Sprintf("%s does not match the email format (email domain).", v.Name)
			}

			e.Add(v.Name, v.Message)
		}
	}
}
