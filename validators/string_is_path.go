package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPath is a validator object
type StringIsPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the path does not exist.
func (v *StringIsPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Field); !os.IsNotExist(err) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must exist", v.Field))
}
