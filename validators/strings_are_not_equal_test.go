package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringsAreNotEqual(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "test", false},
		{"test_fail", "test_true", true},
		{"test with space", " test with space ", false},
		{" test with space second", " test with space second       ", false},
	}

	for _, testCase := range cases {
		v := StringsAreNotEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	v := StringsAreNotEqual{Name: "strings", Field: "test", ComparedField: "test", Message: "Strings match"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Strings match"}, e.Get("strings"))
}

func BenchmarkStringsAreNotEqual_Valid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsAreNotEqual{Name: "strings", Field: " Some string ", ComparedField: " Some string "}
		v.Validate(e)
	}
}

func BenchmarkStringsAreNotEqual_InValid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsAreNotEqual{Name: "strings", Field: " Some string ", ComparedField: " Some string failure"}
		v.Validate(e)
	}
}
