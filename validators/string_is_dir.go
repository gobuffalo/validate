package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsDir is a validator object
type StringIsDir struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a path to directory.
func (v *StringIsDir) Validate(e *validator.Errors) {
	if fi, err := os.Stat(v.Field); !os.IsNotExist(err) {
		if mode := fi.Mode(); mode.IsDir() {
			return
		}
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' is not dir", v.Field))
}
