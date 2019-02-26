package validators

import (
	"fmt"
	"net/url"

	"github.com/s3rj1k/validator"
)

// StringIsURL is a validator object
type StringIsURL struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the field is not a correctly formatted URL.
func (v *StringIsURL) Validate(e *validator.Errors) {

	if v.Field == "http://" || v.Field == "https://" {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s url is empty", v.Name)
		}

		e.Add(v.Name, v.Message)
	}

	parsedURI, err := url.ParseRequestURI(v.Field)
	if err != nil {

		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not match url format, %v", v.Name, err)
		}

		e.Add(v.Name, v.Message)

		return
	}

	if parsedURI.Scheme != "" && parsedURI.Scheme != "http" && parsedURI.Scheme != "https" {

		if v.Message == "" {
			v.Message = fmt.Sprintf("%s invalid url scheme", v.Name)
		}

		e.Add(v.Name, v.Message)
	}
}
