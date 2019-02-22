package validators

import (
	"fmt"
	"net/url"

	"github.com/s3rj1k/validator"
)

// URLIsPresent is a validator object
type URLIsPresent struct {
	Name    string
	Field   string
	Message string
}

// Validate performs the validation to check if URL is formatted correctly
// uses net/url ParseRequestURI to check validity
func (v *URLIsPresent) Validate(e *validator.Errors) {

	if v.Field == "http://" || v.Field == "https://" {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s url is empty", v.Name)
		}

		e.Add(v.Name, v.Message)
	}

	parsedURI, err := url.ParseRequestURI(v.Field)
	if err != nil {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not match url format. Err: %s", v.Name,
				err)
		}

		e.Add(v.Name, v.Message)
	} else {
		if parsedURI.Scheme != "" && parsedURI.Scheme != "http" && parsedURI.Scheme != "https" {
			if v.Message == "" {
				v.Message = fmt.Sprintf("%s invalid url scheme", v.Name)
			}

			e.Add(v.Name, v.Message)
		}
	}
}
