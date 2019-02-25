package validators

import (
	"encoding/json"
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsJSON is a validator object
type StringIsJSON struct {
	Name    string
	Field   string
	Message string
}

// Validate if the string is valid JSON (note: uses json.Unmarshal).
func (v *StringIsJSON) Validate(e *validator.Errors) {

	var js json.RawMessage

	// successful unmarshalling is good
	err := json.Unmarshal([]byte(v.Field), &js)
	if err == nil {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s must be a valid JSON", v.Name))
}
