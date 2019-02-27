package validators

import (
	"fmt"
	"net/url"

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
