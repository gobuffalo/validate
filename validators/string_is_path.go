package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPath is a validator object
type StringIsPath struct {
	Name  string
	Field string
}

// Validate adds an error if the path does not exist.
func (v *StringIsPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Field); !os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must exist", v.Field))
}
