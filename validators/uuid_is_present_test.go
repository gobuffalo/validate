package validators

import (
	"testing"

	u "github.com/gofrs/uuid"
	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_UUIDIsPresent(t *testing.T) {

	r := require.New(t)

	id, err := u.NewV4()
	r.NoError(err)

	v := UUIDIsPresent{Name: "Name", Field: id}
	e := validator.NewErrors()
	v.Validate(e)

	r.Equal(0, e.Count())

	v = UUIDIsPresent{Name: "Name", Field: u.UUID{}}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Name can not be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = UUIDIsPresent{Name: "Name", Field: u.UUID{}, Message: "Field can't be blank"}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))

	e = validator.NewErrors()
	v = UUIDIsPresent{"Name", u.UUID{}, "Field can't be blank"}
	v.Validate(e)

	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("Name"))
}
