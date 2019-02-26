package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// IsDir is a validator object
type IsDir struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if path exists and is not dir.
func (v *IsDir) Validate(e *validator.Errors) {
	if fi, err := os.Stat(v.Field); !os.IsNotExist(err) {
		if mode := fi.Mode(); mode.IsDir() {
			return
		}
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' is not dir", v.Field))
}
