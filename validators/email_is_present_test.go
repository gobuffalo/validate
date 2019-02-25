package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_EmailIsPresent(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		email string
		valid bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}

	for _, test := range tests {
		v := EmailIsPresent{Name: "email", Field: test.email}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equal(test.valid, !e.HasAny())
		if !test.valid {
			r.Equal([]string{"email does not match the email format"}, e.Get("email"))
		}
	}

	v := EmailIsPresent{Name: "email", Field: "", Message: "Email don't match the right format"}
	e := validator.NewErrors()
	v.Validate(e)

	r.Equal(e.Count(), 1)
	r.Equal([]string{"Email don't match the right format"}, e.Get("email"))
}

func Test_EmailLike(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		email string
		valid bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}

	for _, test := range tests {
		v := EmailLike{Name: "email", Field: test.email}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equal(test.valid, !e.HasAny(), test.email)
		if !test.valid {
			r.Equal(e.Get("email"), []string{"email does not match the email format"})
		}
	}

	v := EmailLike{Name: "email", Field: "foo@bar"}
	e := validator.NewErrors()
	v.Validate(e)

	r.Equal(e.Count(), 1)
	r.Equal(e.Get("email"), []string{"email does not match the email format (email domain)"})

	v = EmailLike{Name: "email", Field: "", Message: "Email don't match the right format"}
	e = validator.NewErrors()
	v.Validate(e)

	r.Equal(e.Count(), 1)
	r.Equal(e.Get("email"), []string{"Email don't match the right format"})
}

func BenchmarkEmailIsPresent_IsValid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := EmailLike{Name: "email", Field: "email@gmail.com"}
		v.Validate(e)
	}
}

func BenchmarkEmailLike_IsValid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := EmailIsPresent{Name: "email", Field: "email@gmail.com"}
		v.Validate(e)
	}
}
