package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// FuncValidator is a validator object.
type FuncValidator struct {
	Fn    func() bool
	Name  string
	Field string
}

// Validate is a validation method wrapper.
func (f *FuncValidator) Validate(e *validator.Errors) {
	// for backwards compatibility
	if strings.TrimSpace(f.Name) == "" {
		f.Name = f.Field
	}

	if f.Fn() {
		return
	}

	e.Add(f.Name, fmt.Sprintf("%s result is false", f.Name))
}
