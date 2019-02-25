package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

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
			v.Message = fmt.Sprintf("%s does not match the email format", v.Name)
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
			v.Message = fmt.Sprintf("%s does not match the email format", v.Name)
		}

		e.Add(v.Name, v.Message)
	} else if len(parts) == 2 {
		domain := parts[1]
		// Check that domain is valid
		if len(strings.Split(domain, ".")) < 2 {
			if v.Message == "" {
				v.Message = fmt.Sprintf("%s does not match the email format (email domain)", v.Name)
			}

			e.Add(v.Name, v.Message)
		}
	}
}
