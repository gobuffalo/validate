package validators

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/s3rj1k/validator"
)

// UUIDIsPresent is a validator object
type UUIDIsPresent struct {
	Name    string
	Field   uuid.UUID
	Message string
}

// Validate adds an error if the field is not a valid uuid.
func (v *UUIDIsPresent) Validate(e *validator.Errors) {
	s := v.Field.String()
	if strings.TrimSpace(s) != "" && v.Field != uuid.Nil {
		return
	}

	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	e.Add(v.Name, fmt.Sprintf("%s can not be blank.", v.Name))
}
