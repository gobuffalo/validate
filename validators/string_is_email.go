package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsEmail is a validator object.
type StringIsEmail struct {
	Name  string
	Field string
}

// Validate adds an error if the Field does not match email regexp. See Email const.
func (v *StringIsEmail) Validate(e *validator.Errors) {
	if rxEmail.Match([]byte(v.Field)) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s does not match the email format", v.Name))

}

// SetField sets validator field.
func (v *StringIsEmail) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsEmail) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// StringIsEmailLike is a validator object.
type StringIsEmailLike struct {
	Name  string
	Field string
}

// Validate adds an error if the Field does not correspond to "username@domain" structure.
// It also checks that domain has a domain zone (but does not check if the zone is valid).
// Also allows inner and outer whitespaces.
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

// SetField sets validator field.
func (v *StringIsEmailLike) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsEmailLike) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
