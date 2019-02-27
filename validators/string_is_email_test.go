package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsEmail(t *testing.T) {

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
		v := StringIsEmail{Name: "email", Field: test.email}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(test.valid, !e.HasAny())
		if !test.valid {
			r.Equal([]string{"email does not match the email format"}, e.Get("email"))
		}
	}

	v := StringIsEmail{Name: "email", Field: ""}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(e.Count(), 1)
	r.Equal([]string{"email does not match the email format"}, e.Get("email"))
}

func Test_StringIsEmailLike(t *testing.T) {

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
		v := StringIsEmailLike{Name: "email", Field: test.email}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(test.valid, !e.HasAny(), test.email)
		if !test.valid {
			r.Equal(e.Get("email"), []string{"email does not match the email format"})
		}
	}

	v := StringIsEmailLike{Name: "email", Field: "foo@bar"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(e.Count(), 1)
	r.Equal(e.Get("email"), []string{"email does not match the email format (email domain)"})

	v = StringIsEmailLike{Name: "email", Field: ""}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(e.Count(), 1)
	r.Equal(e.Get("email"), []string{"email does not match the email format"})
}

func BenchmarkStringIsEmailLike(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringIsEmailLike{Name: "email", Field: "email@gmail.com"}
		v.Validate(e)
	}
}

func BenchmarkStringIsEmail(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringIsEmail{Name: "email", Field: "email@gmail.com"}
		v.Validate(e)
	}
}
