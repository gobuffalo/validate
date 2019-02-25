package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// IsDir is a validator object
type IsDir struct {
	Name    string
	Path    string
	Message string
}

// Validate adds an error if path exists and is not dir.
func (v *IsDir) Validate(e *validator.Errors) {

	if fi, err := os.Stat(v.Path); !os.IsNotExist(err) {
		switch mode := fi.Mode(); {
		case mode.IsDir():
			return
		}
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' is not dir", v.Path))
}
