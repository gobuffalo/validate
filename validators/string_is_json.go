package validators

import (
	"encoding/json"
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsJSON is a validator object.
type StringIsJSON struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a valid JSON.
func (v *StringIsJSON) Validate(e *validator.Errors) {

	var js json.RawMessage

	// successful unmarshalling is good
	err := json.Unmarshal([]byte(v.Field), &js)
	if err == nil {
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s must be a valid JSON", v.Name))
}
