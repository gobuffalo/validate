package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// IsNoPath is a validator object
type IsNoPath struct {
	Name    string
	Path    string
	Message string
}

// Validate adds an error if path exists.
func (v *IsNoPath) Validate(e *validator.Errors) {

	if _, err := os.Stat(v.Path); os.IsNotExist(err) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must not exist", v.Path))
}
