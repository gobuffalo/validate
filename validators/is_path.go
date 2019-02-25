package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// IsPath is a validator object
type IsPath struct {
	Name    string
	Path    string
	Message string
}

// Validate adds an error if the path does not exist.
func (v *IsPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Path); !os.IsNotExist(err) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must exist", v.Path))
}
