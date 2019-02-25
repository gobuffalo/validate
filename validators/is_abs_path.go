package validators

import (
	"fmt"
	"path/filepath"

	"github.com/s3rj1k/validator"
)

// IsAbsPath is a validator object
type IsAbsPath struct {
	Name    string
	Path    string
	Message string
}

// Validate adds an error if path is not absolute.
func (v *IsAbsPath) Validate(e *validator.Errors) {

	if v.Path == filepath.Clean(v.Path) && filepath.IsAbs(v.Path) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must be absolute", v.Path))
}
