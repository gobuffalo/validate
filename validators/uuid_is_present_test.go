package validators

import (
	"testing"

	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func Test_UUIDIsPresent(t *testing.T) {
	r := require.New(t)

	id, err := uuid.NewV4()
	r.NoError(err)
	v := UUIDIsPresent{Name: "Name", Field: id}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = UUIDIsPresent{Name: "Name", Field: uuid.UUID{}}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Name can not be blank."})

	errors = validate.NewErrors()
	v = UUIDIsPresent{Name: "Name", Field: uuid.UUID{}, Message: "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})
}
