package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// FuncValidator is a validator object
type FuncValidator struct {
	Fn      func() bool
	Field   string
	Name    string
	Message string
}

// Validate is a validation method wrapper
func (f *FuncValidator) Validate(e *validator.Errors) {
	// for backwards compatibility
	if strings.TrimSpace(f.Name) == "" {
		f.Name = f.Field
	}

	if !f.Fn() {
		e.Add(f.Name, fmt.Sprintf(f.Message, f.Field))
	}
}
