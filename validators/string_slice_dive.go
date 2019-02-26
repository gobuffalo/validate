package validators

import (
	"github.com/s3rj1k/validator"
)

// StringValidator is an interface for string validator objects.
type StringValidator interface {
	Validate(*validator.Errors)
	SetField(s string)
	SetNameIndex(i int)
}

// StringSliceDive is a validator object
type StringSliceDive struct {
	Validator StringValidator
	Field     []string
}

// Validate adds an error if the field is one of the values from the blacklist.
func (v *StringSliceDive) Validate(e *validator.Errors) {
	for i, f := range v.Field {
		v.Validator.SetField(f)
		v.Validator.SetNameIndex(i)
		v.Validator.Validate(e)
	}
}
