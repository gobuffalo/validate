package validators

import (
	"fmt"
	"path/filepath"

	"github.com/s3rj1k/validator"
)

// StringIsAbsPath is a validator object
type StringIsAbsPath struct {
	Name  string
	Field string
}

// Validate adds an error if path is not absolute.
func (v *StringIsAbsPath) Validate(e *validator.Errors) {
	if v.Field == filepath.Clean(v.Field) && filepath.IsAbs(v.Field) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must be absolute", v.Field))
}
