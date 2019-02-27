package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsEmail is a validator object.
type StringIsEmail struct {
	Name  string
	Field string
}

// Validate adds an error if the field does not match email regexp. See Email const.
func (v *StringIsEmail) Validate(e *validator.Errors) {
	if rxEmail.Match([]byte(v.Field)) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s does not match the email format", v.Name))

}

// StringIsEmailLike is a validator object.
type StringIsEmailLike struct {
	Name  string
	Field string
}

// Validate adds an error if the field does not correspond to "username@domain" structure.
// It also checks that domain has a domain zone (but does not check if the zone is valid).
func (v *StringIsEmailLike) Validate(e *validator.Errors) {
	var validStructure = false
	var domainZonePresent = false

	parts := strings.Split(v.Field, "@")

	if len(parts) == 2 && len(parts[0]) > 0 && len(parts[1]) > 0 {
		validStructure = true
	}

	if len(parts) == 2 {
		domain := parts[1]
		// Check that domain is valid
		if len(strings.Split(domain, ".")) >= 2 {
			domainZonePresent = true
		}
	}

	if !validStructure {
		e.Add(v.Name, fmt.Sprintf("%s does not match the email format", v.Name))
		return
	}

	if !domainZonePresent {
		e.Add(v.Name, fmt.Sprintf("%s does not match the email format (email domain)", v.Name))
		return
	}
}
