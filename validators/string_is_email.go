package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsEmail is a validator object.
type StringIsEmail struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field does not match email regexp. See Email const.
func (v *StringIsEmail) Validate(e *validator.Errors) {
	if !rxEmail.Match([]byte(v.Field)) {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not match the email format", v.Name)
		}

		e.Add(v.Name, v.Message)
	}
}

// StringIsEmailLike is a validator object.
type StringIsEmailLike struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field does not correspond to "username@domain" structure.
// It also checks that domain has a domain zone (but does not check if the zone is valid).
func (v *StringIsEmailLike) Validate(e *validator.Errors) {
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
