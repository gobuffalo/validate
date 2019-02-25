package validators

import "github.com/s3rj1k/validator"

// IsNullString is a validator object
type IsNullString struct {
	Name    string
	String  string
	Message string
}

// isNullString is wrapper func
func isNullString(str string) bool {
	if len(str) == 0 { // nolint: megacheck
		return true
	}

	return false
}

// Validate adds an error if string is not (empty).
func (v *IsNullString) Validate(e *validator.Errors) {
	if isNullString(v.String) {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, "string must be emtpy")
}
