package validators

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_TimeIsPresent(t *testing.T) {

	r := require.New(t)

	v := TimeIsPresent{Name: "CreatedAt", Field: time.Now()}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = TimeIsPresent{Name: "CreatedAt", Field: time.Time{}}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"CreatedAt can not be blank"}, e.Get("CreatedAt"))

	e = validator.NewErrors()
	v = TimeIsPresent{Name: "CreatedAt", Field: time.Time{}, Message: "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("CreatedAt"))

	e = validator.NewErrors()
	v = TimeIsPresent{"CreatedAt", time.Time{}, "Field can't be blank"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Field can't be blank"}, e.Get("CreatedAt"))
}
