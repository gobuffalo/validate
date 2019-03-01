package validators

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsURL is a validator object
type StringIsURL struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a correctly formatted URL.
func (v *StringIsURL) Validate(e *validator.Errors) {

	if v.Field == "http://" || v.Field == "https://" {
		e.Add(v.Name, fmt.Sprintf("%s url is empty", v.Name))
		return
	}

	parsedURI, err := url.ParseRequestURI(v.Field)
	if err != nil {
		e.Add(v.Name, fmt.Sprintf("%s does not match url format, %v", v.Name, err))
		return
	}

	if parsedURI.Scheme != "" && parsedURI.Scheme != "http" && parsedURI.Scheme != "https" {
		e.Add(v.Name, fmt.Sprintf("%s invalid url scheme", v.Name))
		return
	}
}

// SetField sets validator field.
func (v *StringIsURL) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsURL) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
