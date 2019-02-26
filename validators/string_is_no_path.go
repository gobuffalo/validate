package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsNoPath is a validator object
type StringIsNoPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if path exists.
func (v *StringIsNoPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Field); os.IsNotExist(err) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must not exist", v.Field))
}
